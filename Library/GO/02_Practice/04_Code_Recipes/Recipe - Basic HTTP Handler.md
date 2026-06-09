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

# Recipe - Basic HTTP Handler

## Когда использовать

- 

## Код

```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte("ok"))
}
```

## Как подключить в проект

```go

```

## Связанная теория

- [[]]

## Где применял

- [[]]
