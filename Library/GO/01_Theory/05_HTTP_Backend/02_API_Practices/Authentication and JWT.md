---
type: theory
area: theory
status: active
created: 2026-06-06
updated: 2026-06-09
tags:
  - area/theory
  - backend
  - go
---

# Authentication and JWT

## Кратко

Authentication подтверждает, кто пользователь. JWT — формат токена, который может нести claims и подписывается секретом или ключом.

## Зачем нужно

Backend использует auth, чтобы отделять публичные endpoint-ы от защищенных и выполнять действия от имени пользователя.

## Типовой поток

1. Пользователь отправляет login/password.
2. Сервер проверяет учетные данные.
3. Сервер выдает access token.
4. Клиент отправляет token в `Authorization: Bearer ...`.
5. Middleware проверяет token и пропускает запрос дальше.

## Правила

- Пароли хранятся только в виде hash.
- JWT должен иметь expiration.
- Секрет подписи не хранится в коде.
- Middleware проверяет подпись и срок действия.
- Не клади в JWT чувствительные данные.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| Пароль plain text | Упрощение | Использовать password hashing |
| JWT без exp | Удобно в dev | Добавить срок действия |
| Секрет в Git | Быстро протестировали | Env/secret storage |

## Связанные темы

- [[01_Theory/09_Production/JWT]]
- [[01_Theory/09_Production/Password_Hashing]]
- [[01_Theory/09_Production/Secrets]]

