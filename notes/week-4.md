# Неделя 4. Конспект (Тесты, CLI, Стек)

## 1. Тесты в Go

- Тесты пишутся в файлах с суффиксом `_test.go`. Функция начинается с `Test` и принимает `t *testing.T`.
- Запуск: `go test -v` (флаг `-v` показывает подробный вывод по каждому кейсу).
- **Табличные тесты:** это способ писать тесты в виде списка кейсов.

```go id="b4f3r1"
tests := []struct {
    name string
    a, b int
    want int
}{
    {"case 1", 2, 2, 4},
    {"case 2", -1, 1, 0},
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
    })
}
```

- `t.Errorf` — помечает тест как упавший, но продолжает выполнение.
- `t.Fatalf` — останавливает тест сразу.
- `httptest` (для HTTP):
  - `httptest.NewRequest` создает фейковый запрос.
  - `httptest.NewRecorder` — фейковый `ResponseWriter` (запоминает ответ хендлера).

**💻 Синтаксис (httptest):**

```go id="c9h2vd"
func TestPingHandler(t *testing.T) {
    // 1. Создаем фейковый запрос и фейковый ответ
    req := httptest.NewRequest("GET", "/ping", nil)
    rec := httptest.NewRecorder()

    // 2. Вызываем хендлер
    pingHandler(rec, req)

    // 3. Проверяем код ответа
    if rec.Code != http.StatusOK {
        t.Errorf("Ожидали 200, получили %d", rec.Code)
    }
}
```

## 2. CLI-утилита (Чтение CSV)

- Аргументы командной строки: `os.Args` (срез).
  - `os.Args[0]` — имя программы.
  - `os.Args[1]` — первый аргумент.
- Пакет `encoding/csv`: `csv.NewReader(file).ReadAll()` возвращает `[][]string`.
- `os.Open(path)` — открыть файл.
- `defer file.Close()` — закрыть после окончания работы.
- Ошибки в Go обрабатываются через `if err != nil { log.Fatal(err) }`.

**💻 Синтаксис (Чтение CSV):**

```go id="m7k3qp"
// 1. Проверяем аргумент
if len(os.Args) < 2 {
    fmt.Fprintln(os.Stderr, "usage: csv-cli <file.csv>")
    os.Exit(1)
}

path := os.Args[1]

// 2. Открываем файл
file, err := os.Open(path)
if err != nil {
    log.Fatalf("open file: %v", err)
}
defer file.Close()

// 3. Парсим CSV
reader := csv.NewReader(file)
rows, err := reader.ReadAll()
if err != nil {
    log.Fatalf("parse csv: %v", err)
}

// 4. Выводим
for _, row := range rows {
    fmt.Println(row)
}
```

## 3. Стек и Очередь (Алгоритмы)

- Стек (Stack) — принцип LIFO (Last In, First Out). Последний пришел — первый ушел. Реализуется через обычный массив (слайс) и операции `append` / `pop`.
- Очередь (Queue) — принцип FIFO (First In, First Out). Первый пришел — первый ушел.
- Монотонный стек — стек, элементы в котором отсортированы (убывают или возрастают). Используется для задач типа "найти следующий больший элемент" (Daily Temperatures).
- Min Stack — дизайн стека, где помимо основных операций (`push`, `pop`, `top`) нужна операция `getMin()` за O(1). Решается через второй (вспомогательный) стек, который хранит минимум на каждом уровне.

**💻 Синтаксис (Min Stack):**

```python id="p6v8wn"
class MinStack(object):
    def __init__(self):
        self.stack = []
        self.min_stack = []

    def push(self, val):
        self.stack.append(val)

        # В min_stack всегда кладем текущий минимум
        if not self.min_stack or val <= self.min_stack[-1]:
            self.min_stack.append(val)
        else:
            self.min_stack.append(self.min_stack[-1])

    def pop(self):
        self.stack.pop()
        self.min_stack.pop()

    def top(self):
        return self.stack[-1]

    def getMin(self):
        return self.min_stack[-1]
```

## 4. Опыт недели (что было тяжело)

- Сложно давалось понимание зачем нужен `httptest` и как передавать в него хендлеры.
- В монотонном стеке путался с индексами и `pop`.
- Начал разбираться с Docker (зачем нужна изоляция).