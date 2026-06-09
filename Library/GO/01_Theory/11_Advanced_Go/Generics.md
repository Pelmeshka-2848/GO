# Generics

## Что это

Generics позволяют писать функции и типы с параметрами типов.

## Зачем нужно

Generics полезны, когда одна и та же логика работает с разными типами без потери type safety.

## Пример формы

```go
func First[T any](items []T) (T, bool) {
    var zero T
    if len(items) == 0 {
        return zero, false
    }
    return items[0], true
}
```

## Когда использовать

- generic containers;
- helper-функции для slice/map;
- algorithms, не зависящие от конкретного типа;
- устранение реального дублирования.

## Когда не использовать

- ради “красоты”;
- если обычная функция читается проще;
- если business logic становится труднее понимать;
- если нужен runtime-polymorphism, а не type parameter.

## Типичные ошибки

- использовать `any` и терять смысл ограничений;
- усложнять простой код;
- писать generic abstraction до появления повторения;
- путать generics и interfaces.

## Связанные темы

- [[01_Theory/01_Go_Basics/Interfaces]]
- [[01_Theory/99_Aliases/01_Syntax/Generics]]

