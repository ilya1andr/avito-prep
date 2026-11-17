# Неделя 7. Конспект (REST API, Слоистая архитектура, PostgreSQL)

## 1. Слоистая архитектура (Главное)

Мы разделили приложение на 3 слоя, чтобы каждый занимался своим делом:

1. **Handler (`main.go` / `handler/`)** — принимает HTTP, парсит JSON, отдает HTTP. Не знает про SQL.
2. **Service (`service/`)** — хранит бизнес-правила (например, "цена не может быть меньше 0"). Не знает про HTTP и SQL.
3. **Repository (`repository/`)** — общается только с БД (делает INSERT, SELECT). Не знает, откуда пришел запрос.

**Зачем:** Если завтра мы захотим поменять базу данных, мы перепишем только Репозиторий. Handler и Service останутся нетронутыми.

## 2. Как слои связаны (Dependency Injection)

Слои собираются в `main.go` как матрешка:

- Создаем подключение к БД (`conn`).
- Передаем `conn` в `NewListingRepository(conn)`.
- Передаем репозиторий в `NewListingService(repo)`.
- Передаем сервис в `NewListingHandler(svc)`.

Хендлер знает про Сервис, Сервис знает про Репозиторий, Репозиторий знает про Базу.

**💻 Синтаксис (Сборка в `main.go`):**

```go
conn, _ := pgx.Connect(ctx, connString)
repo := repository.NewListingRepository(conn)
svc := service.NewListingService(repo)
h := handler.NewListingHandler(svc)
```

## 3. HTTP-методы и коды ответов (REST)

- **GET /listings** — получить список (код 200 OK).
- **GET /listings/{id}** — получить один (код 200 OK, если найден; 404 Not Found, если нет).
- **POST /listings** — создать (код 201 Created).
- **400 Bad Request** — клиент прислал мусор (невалидный JSON, буквы вместо ID).
- **500 Internal Server Error** — ошибка на сервере (база упала).

**💻 Синтаксис (Регистрация роутов в `main.go`):**

```go
mux := http.NewServeMux()

mux.HandleFunc("GET /listings", h.GetListingsHandler)
mux.HandleFunc("POST /listings", h.CreateListing)
mux.HandleFunc("GET /listings/{id}", h.GetListingByID)
mux.HandleFunc("PUT /listings/{id}", h.UpdateListingHandler)
mux.HandleFunc("DELETE /listings/{id}", h.DeleteListingHandler)
```

## 4. JSON в Go (Шпаргалка)

- **Чтение JSON из запроса:** `json.NewDecoder(r.Body).Decode(&input)`
- **Отправка JSON в ответ:** `json.NewEncoder(w).Encode(data)`
- **Теги структур:** `json:"title"` (чтобы поля в Go были с большой буквы, а в JSON — с маленькой).
- `r.Body` — это поток байтов (труба). Его можно прочитать только один раз! Нельзя вызвать `Decode` дважды.

**💻 Синтаксис (Парсинг и ответ):**

```go
// Чтение
var input struct {
    Title string `json:"title"`
    Price int    `json:"price"`
}

if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
    http.Error(w, "Invalid JSON", http.StatusBadRequest)
    return
}

// Ответ
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(createdListing)
```

## 5. PostgreSQL и Go (pgx)

- **Запись/Изменение:** `conn.Exec(ctx, "INSERT INTO ...", args)` (ничего не возвращает).
- **Чтение списка:** `conn.Query(ctx, "SELECT ...")` → возвращает `rows`. Обязательно `defer rows.Close()`!
- **Чтение одной строки:** `conn.QueryRow(ctx, "SELECT ...").Scan(&var)`.
- **Плейсхолдеры:** В SQL пишем `$1, $2`, а значения передаем отдельными аргументами (защита от SQL-инъекций).

**💻 Синтаксис (Репозиторий):**

```go
// Чтение списка
rows, err := r.db.Query(ctx, "SELECT id, title, price FROM listings")
if err != nil {
    return nil, err
}
defer rows.Close()

for rows.Next() {
    var l Listing
    if err := rows.Scan(&l.ID, &l.Title, &l.Price); err != nil {
        return nil, err
    }
    listings = append(listings, l)
}

// Чтение одной
err := r.db.QueryRow(
    ctx,
    "SELECT id, title, price FROM listings WHERE id = $1",
    id,
).Scan(&l.ID, &l.Title, &l.Price)

// Вставка с возвратом ID
err := r.db.QueryRow(
    ctx,
    "INSERT INTO listings (title, price) VALUES ($1, $2) RETURNING id",
    title,
    price,
).Scan(&l.ID)
```

## 6. Обработка ошибок: `pgx.ErrNoRows`

Когда мы ищем объявление по ID и его нет, `QueryRow.Scan()` возвращает ошибку `pgx.ErrNoRows`.

Мы проверяем её через `errors.Is(err, pgx.ErrNoRows)` и отдаем клиенту 404.

**💻 Синтаксис:**

```go
if errors.Is(err, pgx.ErrNoRows) {
    http.Error(w, "Listing not found", http.StatusNotFound)
    return
}
```

## 7. Переменные окружения (12 Factor App)

Пароли и адреса БД нельзя хардкодить в коде!

Используем `os.Getenv("DB_PASSWORD")` и задаем их в терминале:

```powershell
$env:DB_PASSWORD="secret"
```

- **Важно:** Переменные из терминала применяются только к текущей сессии. Пароль в ОС ничего не проверяет, он просто передается базе, а база решает, совпадает ли он с тем, что записано внутри контейнера.

**💻 Синтаксис (Настройка подключения):**

```go
host := os.Getenv("DB_HOST")
if host == "" {
    host = "localhost"
}

user := os.Getenv("DB_USER")
if user == "" {
    user = "postgres"
}

pass := os.Getenv("DB_PASSWORD")
if pass == "" {
    pass = "secret"
}

dbname := os.Getenv("DB_NAME")
if dbname == "" {
    dbname = "postgres"
}

connString := fmt.Sprintf(
    "postgres://%s:%s@%s:5432/%s",
    user,
    pass,
    host,
    dbname,
)
```

## 8. Где я спотыкался (Личные заметки)

- **Сигнатуры хендлеров:** Долго путался с `func(w http.ResponseWriter, r *http.Request)`. Запомнил: `w` — пишем ответ, `r` — читаем запрос.
- `json.NewEncoder` **vs** `w.Write`: `w.Write` отправляет готовый текст, `json.Encode` сам превращает структуру в JSON.
- `context.Background()`: Нужен для долгих операций (подключение к базе). В хендлерах используем `r.Context()`, чтобы запрос отменялся, если клиент закроет вкладку.
- **Переименование:** Ошибка `h.GetListingHandlerByID undefined` возникла из-за того, что метод переименовали в файле хендлера, но забыли в `main.go`. (На будущее: `Shift+F6` в GoLand).
- `make` **и** `docker`: Команды `make` и `docker start` нужно выполнять, когда находишься в правильной папке или терминале.

## 9. Алгоритмы (Паттерны)

- **3Sum (15):** Сортировка + два указателя. Главное — пропускать дубликаты на лету, а не через `if arr not in ans` (это дает TLE).
- **Product of Array Except Self (238):** Запрет на деление. Решается за 2 прохода (префиксные произведения слева и справа).