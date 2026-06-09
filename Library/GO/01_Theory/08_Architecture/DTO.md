# DTO

## Что это

DTO, Data Transfer Object, — структура для передачи данных через границу слоя или API.

## Зачем нужно

DTO отделяет внешний контракт от внутренней модели. API может отдавать не все поля, переименовывать их, скрывать секреты и принимать отдельную форму входа.

## Примеры DTO

```go
type CreateUserRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type UserResponse struct {
    ID    int    `json:"id"`
    Email string `json:"email"`
}
```

## Где использовать

- HTTP request body;
- HTTP response body;
- сообщения очередей;
- внешние API-клиенты;
- границы между слоями, если модель отличается.

## Правила

- Не отдавай password hash в response DTO.
- Request DTO и response DTO часто разные.
- DTO не обязан совпадать с таблицей БД.
- Валидация входа часто начинается с DTO.
- Не создавай DTO без причины внутри маленькой функции.

## Частые ошибки

- использовать DB model как публичный API response;
- принимать лишние поля от клиента;
- смешивать tags разных форматов в одной структуре без необходимости;
- держать бизнес-методы на DTO.

## Связанная теория

- [[01_Theory/05_HTTP_Backend/02_API_Practices/JSON Handling]]
- [[01_Theory/08_Architecture/Layered Architecture]]
- [[01_Theory/01_Go_Basics/Structs]]

