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

# CORS

## Кратко

CORS — браузерный механизм, который контролирует, может ли frontend с одного origin обращаться к API на другом origin.

## Зачем нужно

Если frontend и backend запущены на разных доменах, портах или протоколах, браузер применяет CORS-политику.

## Основные headers

- `Access-Control-Allow-Origin`;
- `Access-Control-Allow-Methods`;
- `Access-Control-Allow-Headers`;
- `Access-Control-Allow-Credentials`.

## Preflight

Для некоторых запросов браузер сначала отправляет `OPTIONS` request, чтобы проверить разрешения.

## Правила

- Не ставь `*` для credentialed-запросов.
- Разрешай только нужные origins.
- Обрабатывай `OPTIONS`.
- Не путай CORS с серверной авторизацией.
- CORS защищает браузерный сценарий, а не API от всех клиентов.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| CORS считается auth | Непонимание роли | Добавить настоящую авторизацию |
| Не обработан OPTIONS | Браузер шлет preflight | Вернуть нужные headers |
| `*` везде | Быстрый dev fix | Настроить список origins |

## Связанные темы

- [[01_Theory/05_HTTP_Backend/02_API_Practices/Middleware]]
- [[01_Theory/05_HTTP_Backend/02_API_Practices/Authentication and JWT]]

