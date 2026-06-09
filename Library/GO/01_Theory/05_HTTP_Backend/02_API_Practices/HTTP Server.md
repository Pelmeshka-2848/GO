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

# HTTP Server

## Кратко

HTTP-сервер — процесс, который слушает порт, принимает запросы и возвращает ответы через handlers.

## Зачем нужно

- REST API;
- healthcheck;
- webhooks;
- internal service endpoints;
- локальная backend-практика.

## Основная идея

В Go HTTP-сервер лучше воспринимать как транспортный слой. Он не должен знать все бизнес-правила и детали базы данных.

## Правила

- Handler должен быть тонким.
- Конфигурация порта приходит извне.
- Ошибка запуска сервера должна обрабатываться.
- Для production задаются read/write/idle timeouts.
- Healthcheck endpoint должен быть простым.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| Вся логика в handler | Быстро писалось в учебном стиле | Вынести service/repository |
| Нет timeout-ов | В dev всё работало | Настроить `http.Server` |
| Порт зашит в коде | Удобно для первой пробы | Перенести в config |

## Связанные темы

- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/HTTP_Server]]
- [[01_Theory/04_Stdlib/net-http package]]
- [[01_Theory/09_Production/Healthcheck]]

