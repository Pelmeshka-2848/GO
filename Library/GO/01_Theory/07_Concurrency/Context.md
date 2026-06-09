# Context

## Что это

`context.Context` передает сигнал отмены, deadline, timeout и request-scoped values через цепочку вызовов.

## Зачем нужно

Context нужен, чтобы долгие операции могли остановиться, когда запрос отменен, timeout истек или сервис завершает работу.

## Основные варианты

- `context.Background()` — корневой context;
- `context.WithCancel` — ручная отмена;
- `context.WithTimeout` — отмена по timeout;
- `context.WithDeadline` — отмена к конкретному времени.

## Ключевые правила

- Context обычно передается первым аргументом функции.
- Не храни context в структуре без веской причины.
- Всегда вызывай `cancel`, когда создаешь context с cancel/timeout.
- Не используй context как обычный контейнер параметров.
- Пробрасывай context в I/O, HTTP и database-вызовы.

## Частые ошибки

- не вызвать `cancel`;
- игнорировать `ctx.Done()`;
- передавать `nil` вместо context;
- складывать в context бизнес-данные, которые должны быть параметрами;
- создавать новый `Background` внутри цепочки вместо проброса существующего.

## Связанная теория

- [[01_Theory/07_Concurrency/Timeouts]]
- [[01_Theory/90_Rules/Context Usage Rules]]
- [[01_Theory/04_Stdlib/context package]]

