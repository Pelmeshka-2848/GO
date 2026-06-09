# HTTP Client

## Что это

HTTP client отправляет HTTP-запросы к внешним или внутренним сервисам.

В Go базовый клиент доступен через `net/http`.

## Зачем нужно

HTTP client нужен для интеграций: внешние API, микросервисы, webhooks, Telegram Bot API, платежи, геокодинг.

## Ключевые правила

- Устанавливай timeout.
- Закрывай response body.
- Проверяй status code.
- Обрабатывай сетевые ошибки отдельно от ошибок API.
- Не создавай новый клиент на каждый запрос без причины.

## Частые ошибки

- использовать клиент без timeout;
- забыть закрыть `resp.Body`;
- читать body до проверки ошибки запроса;
- считать любой HTTP-ответ успешным;
- не ограничивать retry.

## Связанная теория

- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/External_API]]
- [[01_Theory/07_Concurrency/Timeouts]]
- [[01_Theory/01_Go_Basics/Defer]]

