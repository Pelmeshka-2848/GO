---
type: theory
area: theory
status: active
created: 2026-06-06
updated: 2026-06-09
tags:
  - area/theory
  - go
  - stdlib
---

# encoding/json package

## Кратко

`encoding/json` — стандартный пакет для кодирования Go-значений в JSON и декодирования JSON в Go-структуры.

## Зачем нужно

- принимать JSON в REST API;
- отдавать JSON-ответы;
- читать JSON-конфиги;
- работать с внешними API.

## Основные операции

- `json.Marshal` — Go -> JSON bytes;
- `json.Unmarshal` — JSON bytes -> Go;
- `json.NewEncoder(w).Encode(v)` — потоковая запись JSON;
- `json.NewDecoder(r).Decode(&v)` — потоковое чтение JSON.

## Struct tags

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
```

JSON видит только экспортируемые поля.

## Правила

- Проверяй ошибки encode/decode.
- Используй DTO-структуры для внешнего API.
- Не отдавай внутренние поля случайно.
- Для HTTP body удобно использовать `Decoder`.
- Для response body удобно использовать `Encoder`.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| Пустой JSON | Поля не экспортированы | Начать поля с заглавной буквы |
| Ошибка decode игнорируется | Неверный JSON проходит дальше | Проверять `err` |
| `map[string]any` везде | Теряется типизация | Описать struct |

## Связанные темы

- [[01_Theory/03_CLI_And_Files/JSON]]
- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/JSON_API]]
- [[01_Theory/01_Go_Basics/Structs]]

