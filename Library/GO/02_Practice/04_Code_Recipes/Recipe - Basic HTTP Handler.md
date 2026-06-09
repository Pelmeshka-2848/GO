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

- Когда нужен простой healthcheck или учебный endpoint.
- Когда нужно проверить, что HTTP-сервер запускается и отвечает.
- Когда handler не требует тела запроса и сложной бизнес-логики.

## Код

```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte("ok"))
}
```

## Как подключить в проект

Подключается к router или `http.ServeMux` как обычный HTTP handler. В учебном проекте такой endpoint удобно использовать для проверки запуска сервера.

## Связанная теория

- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/Handlers]]
- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/HTTP_Server]]

## Где применял

- [[02_Practice/01_Labs/01_Basic_Backend_Track/003_HTTP_Server/00_Lab Card]]
- [[02_Practice/01_Labs/01_Basic_Backend_Track/004_REST_CRUD_API/00_Lab Card]]
