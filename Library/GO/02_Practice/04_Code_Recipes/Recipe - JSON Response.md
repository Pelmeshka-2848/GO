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

- Когда handler должен вернуть клиенту структурированный JSON.
- Когда нужно единообразно выставлять `Content-Type`.
- Когда API должен возвращать объект, список или ошибку в JSON-формате.

## Код

```go
func writeJSON(w http.ResponseWriter, status int, data any) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    return json.NewEncoder(w).Encode(data)
}
```

## Как подключить в проект

Обычно выносится в helper рядом с HTTP handlers. Handler подготавливает данные и статус, а helper отвечает только за заголовки и кодирование JSON.

## Связанная теория

- [[01_Theory/05_HTTP_Backend/02_API_Practices/JSON Handling]]
- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/HTTP_Status_Codes]]

## Где применял

- [[02_Practice/01_Labs/01_Basic_Backend_Track/004_REST_CRUD_API/00_Lab Card]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/003_REST_API_Todo_Service/00_Lab Card]]
