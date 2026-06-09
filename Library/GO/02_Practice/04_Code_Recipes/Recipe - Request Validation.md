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

# Recipe - Request Validation

## Когда использовать

- Когда входные данные от клиента нельзя доверять напрямую.
- Когда handler принимает JSON с обязательными полями.
- Когда нужно вернуть понятную ошибку до выполнения бизнес-логики.

## Код

```go
if request.Name == "" {
    return errors.New("name is required")
}
```

## Как подключить в проект

Валидацию можно выполнять после декодирования запроса и до вызова service. В более чистой архитектуре проверка формата остается рядом с transport-слоем, а бизнес-правила — в service.

## Связанная теория

- [[01_Theory/05_HTTP_Backend/02_API_Practices/Request Validation]]
- [[01_Theory/05_HTTP_Backend/02_API_Practices/JSON Handling]]

## Где применял

- [[02_Practice/01_Labs/01_Basic_Backend_Track/004_REST_CRUD_API/00_Lab Card]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/003_REST_API_Todo_Service/00_Lab Card]]
