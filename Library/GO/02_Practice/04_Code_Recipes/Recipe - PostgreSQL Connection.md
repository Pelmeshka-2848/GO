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

# Recipe - PostgreSQL Connection

## Когда использовать

- Когда Go-приложению нужно подключиться к PostgreSQL.
- Когда нужно проверить DSN до выполнения запросов.
- Когда CRUD-лабораторная переходит от in-memory хранения к базе данных.

## Код

```go
db, err := sql.Open("postgres", dsn)
if err != nil {
    return err
}
if err := db.Ping(); err != nil {
    return err
}
```

## Как подключить в проект

Подключение обычно создается при старте приложения и передается в repository. Строку подключения лучше брать из конфигурации или переменных окружения.

## Связанная теория

- [[01_Theory/06_Databases/PostgreSQL]]
- [[01_Theory/06_Databases/Connection_String]]
- [[01_Theory/06_Databases/database_sql]]

## Где применял

- [[02_Practice/01_Labs/01_Basic_Backend_Track/005_PostgreSQL_CRUD/00_Lab Card]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/004_REST_API_With_PostgreSQL/00_Lab Card]]
