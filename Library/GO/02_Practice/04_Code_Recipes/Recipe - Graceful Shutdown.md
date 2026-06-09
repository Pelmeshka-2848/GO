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

# Recipe - Graceful Shutdown

## Когда использовать

- 

## Код

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

return server.Shutdown(ctx)
```

## Как подключить в проект

```go

```

## Связанная теория

- [[]]

## Где применял

- [[]]
