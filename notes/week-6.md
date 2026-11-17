# Неделя 6. Конспект (Агрегаты, Оконные функции, Бинарный поиск)

## 1. SQL: Оконные функции

- Не схлопывают строки в отличие от `GROUP BY`. Добавляют новую колонку с вычислением.
- Синтаксис: `ФУНКЦИЯ() OVER (PARTITION BY ... ORDER BY ...)`
- `ROW_NUMBER()` — нумерация строк (1, 2, 3).
- `RANK()` — ранжирование с пропусками (1, 2, 2, 4).
- `LAG()` / `LEAD()` — посмотреть на предыдущую / следующую строку.
- `SUM() OVER (ORDER BY ...)` — накопительный итог (кумулятивная сумма).

**💻 Синтаксис (SQL):**

```sql
-- Нумерация заказов пользователя по дате
SELECT user_id, amount, created_at,
       ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY created_at) AS order_num
FROM orders;

-- Нарастающий итог (кумулятивная сумма)
SELECT user_id, amount,
       SUM(amount) OVER (ORDER BY created_at) AS running_total
FROM orders;

-- Предыдущий заказ (LAG)
SELECT user_id, amount,
       LAG(amount) OVER (PARTITION BY user_id ORDER BY created_at) AS prev_amount
FROM orders;
```

## 2. SQL: EXPLAIN и Индексы

- `EXPLAIN` — показывает план запроса (не выполняя его).
- `EXPLAIN ANALYZE` — выполняет запрос и показывает реальное время.
- **Seq Scan** — полный перебор таблицы (медленно).
- **Index Scan** — поиск по индексу (быстро).
- **Индекс (B-Tree)** — отдельная структура (дерево), которая ускоряет `SELECT`, но замедляет `INSERT/UPDATE/DELETE`.

**💻 Синтаксис (SQL):**

```sql
-- Посмотреть план
EXPLAIN ANALYZE SELECT * FROM users WHERE age = 30;

-- Создать индекс
CREATE INDEX idx_users_age ON users(age);
```

## 3. Go + PostgreSQL (Агрегаты)

- `conn.Query()` — для `SELECT`. Возвращает объект `rows` (поток данных).
- `rows.Next()` — переход к следующей строке.
- `rows.Scan(&var1, &var2)` — запись значений колонок в переменные по ссылке.
- `rows.Close()` — обязательно закрывать, чтобы освободить ресурсы базы.
- `flag.Bool()` — пакет `flag` для создания флагов командной строки (например, `--stats`).

**💻 Синтаксис (Go):**

```go
// Флаг командной строки (запуск: go run . --stats)
statsFlag := flag.Bool("stats", false, "Show statistics")
flag.Parse()

// Чтение агрегатов
rows, err := conn.Query(ctx, "SELECT city, COUNT(*), AVG(age) FROM people GROUP BY city")
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
    var city string
    var count int
    var avgAge float64

    if err := rows.Scan(&city, &count, &avgAge); err != nil {
        log.Fatal(err)
    }

    fmt.Printf(
        "Город: %s, Людей: %d, Средний возраст: %.1f\n",
        city,
        count,
        avgAge,
    )
}
```

## 4. Алгоритмы: Бинарный поиск

- Работает только на **отсортированных** массивах. Время: O(log n).
- `Search in Rotated Sorted Array (33)`: нужно определять, какая половина массива (левая или правая) точно отсортирована. Сравниваем `nums[l] <= nums[mid]`.
- `Find Minimum in Rotated Sorted Array (153)`: сравниваем `nums[mid]` с правым краем `nums[r]`. Если `nums[mid] > nums[r]`, идем вправо, иначе влево.