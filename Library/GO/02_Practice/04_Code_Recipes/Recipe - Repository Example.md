---
type: recipe
area: practice
status: draft
created: 2026-06-06
updated:
tags:
  - area/practice
  - type/snippet
---

# Recipe - Repository Example

## Когда использовать

- Когда нужно отделить бизнес-логику от конкретного хранилища.
- Когда service должен работать с абстракцией, а не с SQL напрямую.
- Когда проект готовится к тестированию через подмену repository.

## Код

```go
type UserRepository interface {
    FindByID(ctx context.Context, id int64) (*User, error)
}
```

## Как подключить в проект

Repository-интерфейс обычно используется service-слоем. Конкретная реализация может работать с PostgreSQL, in-memory хранилищем или тестовой заглушкой.

## Связанная теория

- [[01_Theory/06_Databases/Repository_Layer]]
- [[01_Theory/08_Architecture/Layered_Architecture]]
- [[01_Theory/90_Rules/Interface Design Rules]]

## Где применял

- [[02_Practice/01_Labs/01_Basic_Backend_Track/005_PostgreSQL_CRUD/00_Lab Card]]
- [[02_Practice/01_Labs/01_Basic_Backend_Track/007_Clean_Architecture_API/00_Lab Card]]
