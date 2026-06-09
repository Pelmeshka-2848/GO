# Swagger OpenAPI

## Что это

OpenAPI — спецификация описания HTTP API. Swagger — набор инструментов вокруг OpenAPI.

## Зачем нужно

Документация API нужна, чтобы клиент понимал endpoint-ы, request/response форматы, статусы ошибок и auth-схему.

## Что описывать

- paths;
- HTTP methods;
- request body;
- response body;
- status codes;
- query/path parameters;
- auth;
- error format.

## Польза

- единый контракт API;
- автогенерация документации;
- удобное ручное тестирование;
- меньше разночтений между frontend и backend;
- основа для client generation.

## Правила

- Документация должна соответствовать реальному API.
- Error response тоже часть контракта.
- DTO должны быть описаны явно.
- Auth-схема должна быть видна.
- Не документируй внутренние endpoint-ы как публичные без причины.

## Частые ошибки

- документация устарела;
- описан только happy path;
- не указаны ошибки;
- request/response не совпадает с кодом;
- нет примеров.

## Связанные темы

- [[01_Theory/05_HTTP_Backend/02_API_Practices/REST API]]
- [[01_Theory/05_HTTP_Backend/02_API_Practices/JSON Handling]]
- [[01_Theory/05_HTTP_Backend/02_API_Practices/Authentication and JWT]]

