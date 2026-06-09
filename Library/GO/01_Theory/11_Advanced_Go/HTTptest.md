# HTTptest

## Что это

`httptest` — стандартный пакет для тестирования HTTP handlers и HTTP-серверов.

## Зачем нужно

Он позволяет проверить handler без настоящего сетевого порта и браузера.

## Основные инструменты

- `httptest.NewRequest` — создать тестовый request;
- `httptest.NewRecorder` — записать response handler-а;
- `httptest.NewServer` — поднять временный HTTP-сервер для тестов client-кода.

## Что проверять

- status code;
- response body;
- headers;
- обработку невалидного запроса;
- auth/middleware поведение;
- JSON response format.

## Правила

- Тест handler-а должен проверять HTTP-контракт.
- Не поднимай реальный порт без необходимости.
- Зависимости handler-а лучше подменять fake-реализациями.
- Проверяй error cases, не только `200 OK`.

## Типичные ошибки

- тестировать handler через реальный localhost;
- не проверять status code;
- сравнивать JSON как строку без необходимости;
- тянуть настоящую БД в unit-тест handler-а.

## Связанные темы

- [[01_Theory/05_HTTP_Backend/02_API_Practices/HTTP Server]]
- [[01_Theory/05_HTTP_Backend/02_API_Practices/JSON Handling]]
- [[01_Theory/11_Advanced_Go/Testing]]


