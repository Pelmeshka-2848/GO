# Middleware

## Что это

Middleware — промежуточный слой обработки HTTP-запроса вокруг handler-а.

## Зачем нужно

Middleware выносит повторяющуюся логику из handler-ов:

- логирование;
- проверка токена;
- CORS;
- recovery после panic;
- rate limiting;
- request ID;
- timeout.

## Модель работы

1. Запрос приходит в middleware.
2. Middleware выполняет свою проверку или подготовку.
3. Middleware передает запрос следующему handler-у.
4. После handler-а middleware может выполнить завершающую логику.

## Ключевые правила

- Middleware должен делать одну понятную вещь.
- Порядок middleware важен.
- Auth middleware не должен содержать бизнес-логику.
- Ошибки middleware должны возвращаться в общем API-формате.
- Не прячь сложное состояние в middleware без необходимости.

## Частые ошибки

- делать один большой middleware “для всего”;
- забыть вызвать следующий handler;
- неправильно расположить auth и logging;
- менять request context без ясного контракта.

## Связанная теория

- [[01_Theory/05_HTTP_Backend/02_API_Practices/Middleware]]
- [[01_Theory/09_Production/Auth]]
- [[01_Theory/08_Architecture/Rate_Limiting]]

