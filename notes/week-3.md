# Неделя 3 — Конспект

## Методы

- На значении (`t T`) — работает с копией, не может менять получателя.
- На указателе (`t *T`) — работает с оригиналом, может менять.
- Если один метод на указателе — делай все на указателе.
- `*T` может вызывать оба вида; `T` — только методы на значении.

**💻 Синтаксис:**

```go
type Counter struct {
    n int
}

func (c Counter) Get() int { return c.n } // На значении

func (c *Counter) Inc() { c.n++ } // На указателе
```

## Интерфейсы

- Набор методов. Реализация неявная — `implements` не пишется.
- nil-интерфейс (тип=nil, значение=nil) ≠ интерфейс с nil-значением (тип задан, значение nil).
- Пустой интерфейс `any` — реализует любой тип.
- Type assertion: `x.(T)`. Type switch: `switch v := x.(type)`.
- Метод на `*T` реализует интерфейс только через `*T`, не через `T`.

**💻 Синтаксис:**

```go
type Shape interface {
    Area() float64
}

// Реализация
type Circle struct {
    r float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.r * c.r
}

// Проверка типа
if circle, ok := shape.(Circle); ok {
    fmt.Println(circle.r)
}
```

## Ошибки

- `error` — интерфейс с методом `Error() string`. Не исключение.
- Возвращается последним значением.
- `errors.New` — простая. `fmt.Errorf("%w")` — с обёрткой.
- `errors.Is` — проверка «это та ошибка?» (по цепочке).
- `errors.As` — извлечение конкретного типа из цепочки.
- Ошибки не игнорируются — компилятор заругается.

**💻 Синтаксис:**

```go
// Создание
err := errors.New("something went wrong")
errWithWrap := fmt.Errorf("db error: %w", err)

// Проверка
if errors.Is(err, ErrNotFound) {
    // ...
}
```

## Вопросы к себе

- Чем метод на указателе отличается от метода на значении?
- Что внутри интерфейса? Почему nil-интерфейс ≠ интерфейс с nil-значением?
- Зачем `errors.Is`, если есть `==`?

## Exercise: Errors

- `error` — встроенный интерфейс с методом `Error() string`.
- Любой тип с этим методом — error. `implements` не нужен.
- Возвращать ошибку последним значением: `(T, error)`.
- Нет ошибки — `return ..., nil`.
- `type ErrNegativeSqrt float64` — новый тип на основе `float64`.
- Метод `Error()` на нём делает его error.
- Внутри `Error()` **НЕ** передавай `e` в `fmt.Sprintf` без преобразования — будет рекурсия. `fmt.Sprintf("%v", float64(e))` — правильно.
- `ErrNegativeSqrt(x)` — преобразование типа `float64` → `ErrNegativeSqrt`.

**💻 Синтаксис (Error Exercise):**

```go
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    // Обязательно float64(e), иначе бесконечная рекурсия
    return fmt.Sprintf(
        "cannot Sqrt negative number: %v",
        float64(e),
    )
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    }

    return math.Sqrt(x), nil
}
```

## Exercise: Readers

- `io.Reader` — интерфейс с одним методом `Read(p []byte) (int, error)`.
- `Read` заполняет буфер `p` и возвращает, сколько байт записал.
- `n = len(p)` — заполнил весь буфер.
- `err = nil` — поток бесконечный, ошибки нет. `io.EOF` — только когда поток закончился.
- `p[i] = 'A'` — одинарные кавычки = rune (символ), двойные = строка.

**💻 Синтаксис (Reader Exercise):**

```go
type MyReader struct{}

func (r MyReader) Read(p []byte) (int, error) {
    for i := range p {
        p[i] = 'A' // Заполняем буфер символом A
    }

    return len(p), nil // Заполнили всё, ошибки нет
}
```