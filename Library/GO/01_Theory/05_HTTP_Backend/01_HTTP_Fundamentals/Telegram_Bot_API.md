# Telegram Bot API

## Что это

Telegram Bot API — HTTP API для управления ботом Telegram: получение сообщений, отправка ответов, работа с командами и webhooks.

## Зачем нужно

Это удобный учебный пример внешнего API: есть токен, HTTP-запросы, JSON-ответы, rate limits и обработка ошибок.

## Основные способы получения обновлений

- long polling: приложение периодически запрашивает новые события;
- webhook: Telegram отправляет события на endpoint приложения.

## Что важно для Go

- токен нельзя хранить в коде;
- HTTP client должен иметь timeout;
- ответы API нужно декодировать из JSON;
- ошибки внешнего API нужно отделять от ошибок своего приложения;
- для webhook нужен публично доступный HTTPS endpoint.

## Ключевые правила

- Для первых лабораторных проще long polling.
- Webhook лучше изучать после HTTP-сервера.
- Не логируй полный bot token.
- Учитывай rate limits.
- Отделяй Telegram-клиент от бизнес-логики бота.

## Частые ошибки

- закоммитить токен бота;
- не проверять ответ Telegram API;
- делать бесконечный polling без паузы;
- смешать обработку команд, HTTP-клиент и хранение состояния в одном файле.

## Связанная теория

- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/External_API]]
- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/API_Tokens]]
- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/Rate_Limits]]

