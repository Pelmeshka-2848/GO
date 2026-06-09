# strings.HasPrefix

## Суть

`strings.HasPrefix` проверяет, начинается ли строка с указанного префикса.

```go
strings.HasPrefix("4000123456789017", "4000")
```

Результат:

```go
true
```

## Сигнатура

```go
func HasPrefix(s, prefix string) bool
```

## Gotcha

Пустой prefix считается началом любой строки.

```go
strings.HasPrefix("40001234", "") // true
```

Если пустой prefix невозможен по бизнес-логике, это нужно проверять отдельно.

## Связанные заметки

- [[01_Theory/99_Atomic_Go/023_range_Loop]]

