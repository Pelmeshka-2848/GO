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

# JSON Handling

## Кратко

JSON handling в backend — это чтение JSON-запросов, валидация данных и формирование JSON-ответов.

## Зачем нужно

REST API чаще всего общается с клиентами через JSON. Ошибка в обработке JSON быстро превращается в плохой API-контракт.

## Основная идея

Входной JSON лучше декодировать в request DTO, а ответ формировать из response DTO.

## Правила

- Проверяй ошибку декодирования.
- Не используй внутреннюю модель БД как публичный ответ без причины.
- Выставляй `Content-Type: application/json`.
- Возвращай единый формат ошибок.
- Ограничивай размер body там, где API публичный.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| Поля не попали в JSON | Они не экспортированы | Использовать заглавные поля и tags |
| Игнор decode error | Спешка | Вернуть `400 Bad Request` |
| Разные error response | Нет helper-а | Сделать единый формат |

## Связанные темы

- [[01_Theory/04_Stdlib/encoding-json package]]
- [[01_Theory/03_CLI_And_Files/JSON]]
- [[01_Theory/05_HTTP_Backend/02_API_Practices/Request Validation]]

