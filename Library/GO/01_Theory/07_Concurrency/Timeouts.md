# Timeouts

## Что это

Timeout — ограничение времени выполнения операции.

## Зачем нужно

Timeout защищает программу от бесконечного ожидания: сетевого запроса, ответа worker-а, блокировки channel, запроса к базе.

## Где нужен timeout

- HTTP client;
- HTTP server;
- запрос к базе данных;
- внешний API;
- worker task;
- ожидание результата через channel.

## Основные подходы

- `context.WithTimeout`;
- `time.After`;
- timeout-настройки `http.Client`;
- timeout-настройки `http.Server`.

## Ключевые правила

- Внешние вызовы должны иметь timeout.
- Не используй один глобальный timeout для всех случаев без смысла.
- После `context.WithTimeout` вызывай `cancel`.
- Отличай timeout от обычной бизнес-ошибки.
- Логируй, какая операция превысила время.

## Частые ошибки

- HTTP client без timeout;
- забыть `cancel()`;
- использовать слишком маленький timeout и получать случайные падения;
- не пробрасывать context ниже по слоям.

## Связанная теория

- [[01_Theory/07_Concurrency/Context]]
- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/HTTP_Client]]
- [[01_Theory/09_Production/Observability]]

