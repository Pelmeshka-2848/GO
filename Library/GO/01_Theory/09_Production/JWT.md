# JWT

## Что это

JWT, JSON Web Token, — компактный подписанный token с набором claims.

## Зачем нужно

JWT часто используют как access token для stateless authentication в HTTP API.

## Структура

JWT состоит из трех частей:

- header;
- payload;
- signature.

Payload не является секретным сам по себе. Его можно прочитать, если token попал к клиенту.

## Важные claims

- `sub` — subject, обычно user id;
- `exp` — expiration time;
- `iat` — issued at;
- `iss` — issuer;
- `aud` — audience.

## Правила

- JWT должен иметь срок действия.
- Секрет подписи не хранится в коде.
- Проверяй signature и expiration.
- Не клади чувствительные данные в payload.
- Для refresh token нужна отдельная стратегия.

## Частые ошибки

- считать JWT зашифрованным;
- не проверять `exp`;
- использовать слабый secret;
- хранить слишком много данных в claims;
- не иметь механизма отзыва при необходимости.

## Связанные темы

- [[01_Theory/05_HTTP_Backend/02_API_Practices/Authentication and JWT]]
- [[01_Theory/09_Production/Secrets]]

