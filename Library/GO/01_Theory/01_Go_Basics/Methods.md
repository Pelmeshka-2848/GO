# Methods

## Что это

Метод — это функция, привязанная к конкретному типу через receiver.

```go
type User struct {
    Name string
}

func (u User) DisplayName() string {
    return u.Name
}
```

## Зачем нужно

Методы помогают связать поведение с типом: пользователь может менять имя, задача может завершаться, счет может пересчитывать итог.

## Value receiver и pointer receiver

Value receiver получает копию значения:

```go
func (u User) DisplayName() string {
    return u.Name
}
```

Pointer receiver получает адрес значения и может менять исходную структуру:

```go
func (u *User) Rename(name string) {
    u.Name = name
}
```

## Ключевые правила

- Если метод меняет состояние, используй pointer receiver.
- Если структура большая, pointer receiver может избежать лишнего копирования.
- Не смешивай value и pointer receiver без причины для одного типа.
- Метод должен быть связан со смыслом типа, а не просто быть случайной функцией.

## Частые ошибки

- ожидать изменения структуры при value receiver;
- привязывать к типу методы, которые не используют receiver;
- делать методы слишком большими;
- использовать pointer receiver без проверки `nil`, если nil возможен.

## Связанная теория

- [[01_Theory/01_Go_Basics/Structs]]
- [[01_Theory/01_Go_Basics/Pointers]]
- [[01_Theory/99_Atomic_Go/020_Pointers]]

