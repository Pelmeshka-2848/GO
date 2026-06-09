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

# Middleware

## Кратко

Middleware — обертка вокруг HTTP-handler-а, которая добавляет общую обработку запроса.

## Зачем нужно

- логирование;
- request ID;
- CORS;
- auth;
- recovery;
- rate limiting;
- timeout.

## Основная идея

Middleware должен делать одну инфраструктурную задачу и передавать управление следующему handler-у.

## Правила

- Порядок middleware важен.
- Middleware не должен становиться местом бизнес-логики.
- Auth middleware проверяет доступ, но не выполняет сценарий пользователя.
- Ошибки middleware должны использовать общий формат API.
- Данные в context должны иметь понятный контракт.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| Не вызван next | Забыли передать управление | Вызвать следующий handler |
| Один middleware для всего | Быстро росло | Разделить по задачам |
| Business logic внутри | Удобно добраться до request | Перенести в service |

## Связанные темы

- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/Middleware]]
- [[01_Theory/05_HTTP_Backend/02_API_Practices/Authentication and JWT]]
- [[01_Theory/05_HTTP_Backend/02_API_Practices/CORS]]

