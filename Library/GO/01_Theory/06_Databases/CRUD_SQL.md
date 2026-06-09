# CRUD SQL

## Что это

CRUD — четыре базовые операции с данными:

- Create;
- Read;
- Update;
- Delete.

В SQL им обычно соответствуют `INSERT`, `SELECT`, `UPDATE`, `DELETE`.

## Зачем нужно

Почти любой backend API работает с CRUD-операциями: создать пользователя, получить список задач, обновить статус, удалить запись.

## Базовая модель

| Операция | SQL | HTTP в REST |
|---|---|---|
| Create | `INSERT` | `POST` |
| Read | `SELECT` | `GET` |
| Update | `UPDATE` | `PUT` или `PATCH` |
| Delete | `DELETE` | `DELETE` |

## Ключевые правила

- `SELECT` должен возвращать только нужные поля.
- `UPDATE` и `DELETE` почти всегда должны иметь `WHERE`.
- Входные значения передаются параметрами, а не склейкой строк.
- После `INSERT` часто нужно вернуть созданный `id`.
- Ошибку “запись не найдена” нужно отличать от внутренней ошибки БД.

## Частые ошибки

- забыть `WHERE` в `UPDATE` или `DELETE`;
- использовать `SELECT *` в API-слое без необходимости;
- не проверять количество затронутых строк;
- не обрабатывать конфликт уникальности;
- смешивать SQL прямо в HTTP-handler-е.

## Связанная теория

- [[01_Theory/06_Databases/SQL_Injection]]
- [[01_Theory/06_Databases/Repository_Layer]]
- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/REST_API]]

