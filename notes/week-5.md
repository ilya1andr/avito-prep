# Неделя 5. Конспект (SQL, PostgreSQL, Строки)

## 1. SQL: JOIN

- **INNER JOIN**: показывает только те строки, где есть совпадения в обеих таблицах (пересечение).
- **LEFT JOIN**: показывает ВСЕ строки из левой таблицы + совпадения из правой. Если пары нет — `NULL`.
- **RIGHT JOIN**: то же самое, но наоборот (все из правой).
- **FULL JOIN**: все строки из обеих таблиц.
- **Правило**: всегда указывать условие соединения через `ON` (например, `ON users.id = orders.user_id`), иначе база сделает декартово произведение (`CROSS JOIN`).

**💻 Синтаксис (SQL):**

```sql
-- INNER JOIN
SELECT u.name, o.amount
FROM users u
INNER JOIN orders o ON u.id = o.user_id;

-- LEFT JOIN
SELECT u.name, o.amount
FROM users u
LEFT JOIN orders o ON u.id = o.user_id;
```

## 2. SQL: Агрегаты и группировка

- Агрегаты: `COUNT`, `SUM`, `AVG`, `MIN`, `MAX`.
- `GROUP BY`: схлопывает строки в группы по колонке.
- `HAVING`: фильтр для результатов агрегации (пишется **ПОСЛЕ `GROUP BY`**). В отличие от `WHERE`, который фильтрует строки **ДО группировки**.
- **Правило**: всё, что в `SELECT` не обернуто в агрегат, должно быть в `GROUP BY`.
- `COUNT(DISTINCT колонка)`: считает только уникальные значения (например, уникальных клиентов).

**💻 Синтаксис (SQL):**

```sql
-- Группировка и HAVING
SELECT city, COUNT(*) AS user_count
FROM users
WHERE age > 18        -- Фильтр ДО группировки
GROUP BY city
HAVING COUNT(*) > 10; -- Фильтр ПОСЛЕ группировки

-- Уникальные клиенты
SELECT COUNT(DISTINCT user_id)
FROM orders;
```

## 3. SQL: Подзапросы и CTE

- **Подзапрос**: запрос внутри запроса в скобках `()`.
- В `WHERE`: для фильтрации (например, `WHERE id IN (SELECT user_id ...)`).
- В `FROM`: как временная таблица (обязательно давать ей алиас).
- **CTE (Common Table Expression)**: конструкция `WITH имя AS (SELECT ...)`.
- Позволяет вынести подзапросы наверх, чтобы читать код сверху вниз.
- Если шагов больше двух — всегда лучше использовать CTE, а не матрешку из скобок.

**💻 Синтаксис (SQL):**

```sql
-- CTE
WITH UserRevenue AS (
    SELECT user_id, SUM(amount) AS total
    FROM orders
    GROUP BY user_id
)
SELECT *
FROM UserRevenue
WHERE total > 1000;
```

## 4. SQL: Работа с датами

- `DATE_TRUNC('month', date)`: обрезает дату до первого числа месяца (возвращает тип date/timestamp).
- `EXTRACT(MONTH FROM date)`: достает номер месяца, дня (`DAY`) или недели (`WEEK`) из даты.

**💻 Синтаксис (SQL):**

```sql
SELECT DATE_TRUNC('month', created_at) AS month, AVG(amount)
FROM orders
GROUP BY month;

SELECT *
FROM orders
WHERE EXTRACT(MONTH FROM created_at) = 1;
```

## 5. Go + PostgreSQL

- Для работы с базой используем драйвер `pgx` (пакет `github.com/jackc/pgx/v5`).
- Подключение: `pgx.Connect(ctx, connString)`.
- Выполнение запроса без получения строк (например, INSERT): `conn.Exec(ctx, sql, args...)`.
- **Плейсхолдеры**: в SQL-запросе пишем `$1, $2, $3`, а значения передаем отдельными аргументами. Это защищает от SQL-инъекций.
- **Переменные окружения**: пароли и адреса баз НЕЛЬЗЯ хардкодить. Их читают через `os.Getenv("DB_PASSWORD")`.

**💻 Синтаксис (Go):**

```go
conn, err := pgx.Connect(ctx, "postgres://postgres:secret@localhost:5432/postgres")

// INSERT
_, err = conn.Exec(
    ctx,
    "INSERT INTO people (name, age) VALUES ($1, $2)",
    "Alice",
    30,
)

// Переменные окружения
pass := os.Getenv("DB_PASSWORD")
```

## 6. Docker

- **Контейнер** — это изолированная среда (капсула) для запуска программы.
- Запуск базы: `docker run --name pg -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:16`.
- Остановить: `docker stop pg`.
- Зайти внутрь контейнера: `docker exec -it pg bash`.
- Посмотреть логи: `docker logs pg`.

**💻 Синтаксис (Терминал):**

```powershell
# Запустить контейнер заново
docker start pg

# Зайти в psql (базу данных)
docker exec -it pg psql -U postgres

# Внутри psql посмотреть таблицы
\dt

# Выйти
\q
```

## 7. Алгоритмы: Строки

- **Скользящее окно (Sliding Window)**: для поиска подстрок. Два указателя (`left`, `right`), хеш-таблица для символов. (Задача: Longest Substring Without Repeating Characters).
- **Расширение от центра (Expand Around Center)**: для палиндромов. Идем по каждому символу, считаем его центром и расширяемся влево-вправо. Не забываем про четные палиндромы (центр между буквами). (Задача: Longest Palindromic Substring).
- **Парсинг строк (Симуляция)**: пошаговая обработка строки с проверкой границ и знаков. (Задача: String to Integer).