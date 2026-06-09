---
type: rule
area: theory
status: active
created: 2026-06-06
updated: 2026-06-09
tags:
  - area/theory
  - type/rule
---

# Interface Design Rules

## Правило

Интерфейс должен описывать поведение, которое нужно потребителю. Не создавай interface только потому, что есть struct.

## Почему так

В Go интерфейсы реализуются неявно. Поэтому лучше объявлять маленький interface рядом с кодом, которому нужно поведение, а не заранее строить абстракции “на будущее”.

## Хорошо

```go
type UserFinder interface {
    FindUser(id int) (User, error)
}
```

Такой interface говорит: этому коду нужно только находить пользователя.

## Плохо

```go
type UserRepository interface {
    Create(User) error
    Update(User) error
    Delete(int) error
    Find(int) (User, error)
    List() ([]User, error)
    Count() (int, error)
}
```

Плохо, если потребителю нужен только один метод, а interface заставляет реализовывать все.

## Когда можно нарушить

- На границе крупного слоя, где полный контракт действительно нужен.
- В generated-коде.
- В учебной заметке, где нужно показать общий набор операций.

## Связанные темы

- [[01_Theory/01_Go_Basics/Interfaces]]
- [[01_Theory/91_Lifehacks/How to Think About Interfaces]]
- [[01_Theory/08_Architecture/Dependency Injection]]

