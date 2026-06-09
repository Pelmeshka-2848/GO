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

# Recipe - Middleware

## Когда использовать

- Когда одну проверку нужно применить к нескольким handlers.
- Когда нужно добавить logging, auth, recover или CORS.
- Когда важно оставить handler сфокусированным на бизнес-действии.

## Код

```go
func middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        next.ServeHTTP(w, r)
    })
}
```

## Как подключить в проект

Middleware оборачивает следующий `http.Handler`. Его можно подключать к отдельному маршруту, группе маршрутов или ко всему router.

## Связанная теория

- [[01_Theory/05_HTTP_Backend/02_API_Practices/Middleware]]
- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/Handlers]]

## Где применял

- [[02_Practice/01_Labs/01_Basic_Backend_Track/006_Auth_JWT/00_Lab Card]]
- [[02_Practice/02_Development_Cases/Case - Middleware Chain]]
