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

# Recipe - JSON Response

## Когда использовать

- 

## Код

```go
func writeJSON(w http.ResponseWriter, status int, data any) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    return json.NewEncoder(w).Encode(data)
}
```

## Как подключить в проект

```go

```

## Связанная теория

- [[]]

## Где применял

- [[]]
