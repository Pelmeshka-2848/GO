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

# net/http package

## Кратко

`net/http` — стандартный пакет Go для HTTP-серверов и HTTP-клиентов.

## Зачем нужно

- запускать REST API;
- писать handlers;
- принимать HTTP-запросы;
- возвращать HTTP-ответы;
- обращаться к внешним API.

## Основные серверные элементы

- `http.Handler`;
- `http.HandlerFunc`;
- `http.ResponseWriter`;
- `*http.Request`;
- `http.Server`;
- `http.ServeMux`.

## Основные клиентские элементы

- `http.Client`;
- `http.NewRequestWithContext`;
- `http.Response`;
- `http.Transport`.

## Правила

- На server side проверяй ошибку запуска сервера.
- На client side задавай timeout.
- Закрывай `resp.Body` после успешного ответа клиента.
- Для JSON выставляй `Content-Type`.
- Для production-сервера настраивай timeouts.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| Client без timeout | Запрос может зависнуть | Настроить `http.Client{Timeout: ...}` |
| Не закрыт body | Утечка ресурсов | `defer resp.Body.Close()` |
| Handler делает всё | Смешаны слои | Вынести service/repository |

## Связанные темы

- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/HTTP_Server]]
- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/Handlers]]
- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/HTTP_Client]]

