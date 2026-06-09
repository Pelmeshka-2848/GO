---
type: theory
area: theory
status: active
created: 2026-06-06
updated: 2026-06-09
tags:
  - area/theory
  - backend
  - go
---

# REST API

## Кратко

REST API — стиль проектирования HTTP API, где URL представляет ресурс, а HTTP-метод задает действие.

## Зачем нужно

- строить понятные API-контракты;
- отделять действия от путей;
- использовать стандартные HTTP-методы и статусы;
- делать API предсказуемым для клиента.

## Основная идея

Путь отвечает на вопрос “с каким ресурсом работаем”, метод — “что делаем”.

```text
GET    /tasks
POST   /tasks
GET    /tasks/{id}
PATCH  /tasks/{id}
DELETE /tasks/{id}
```

## Правила

- Используй существительные в URL.
- Не пиши действия в пути, если их выражает HTTP-метод.
- Возвращай корректные status codes.
- Ошибки API должны иметь единый формат.
- Не раскрывай внутреннюю структуру БД через API.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| `/getUsers` | RPC-мышление | `GET /users` |
| Всегда `200` | Упрощение | Подбирать статус по результату |
| SQL-модель наружу | Нет DTO | Разделить internal model и response |

## Связанные темы

- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/REST_API]]
- [[01_Theory/05_HTTP_Backend/02_API_Practices/JSON Handling]]
- [[01_Theory/05_HTTP_Backend/02_API_Practices/Request Validation]]

