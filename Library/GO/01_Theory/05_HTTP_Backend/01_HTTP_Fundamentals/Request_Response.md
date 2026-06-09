# Request Response

## Что это

HTTP-запрос описывает, что клиент хочет сделать. HTTP-ответ сообщает результат обработки.

## Request

Запрос содержит:

- method: `GET`, `POST`, `PUT`, `DELETE`;
- path: например `/tasks/1`;
- query parameters;
- headers;
- body.

## Response

Ответ содержит:

- status code;
- headers;
- body.

Для JSON API тело ответа обычно содержит JSON.

## Ключевые правила

- Не читай body без необходимости.
- Проверяй ошибки декодирования.
- Выставляй `Content-Type` для JSON-ответов.
- Возвращай статус, соответствующий результату операции.
- Не возвращай внутренние технические детали пользователю.

## Частые ошибки

- всегда возвращать `200 OK`;
- забыть закрыть body у исходящих HTTP-запросов;
- не проверять Content-Type там, где это важно;
- смешать query parameters и path parameters.

## Связанная теория

- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/HTTP_Status_Codes]]
- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/JSON_API]]

