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

- 

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

```go

```

## Связанная теория

- [[]]

## Где применял

- [[]]
