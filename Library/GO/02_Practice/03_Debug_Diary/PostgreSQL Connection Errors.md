---
type: debug
area: practice
status: active
created: 2026-06-06
updated:
tags:
  - area/practice
  - type/error
---

# PostgreSQL Connection Errors

## Симптомы

```text
connection refused
password authentication failed
database does not exist
role does not exist
relation does not exist
```

## Где возникло

- проект: REST API с PostgreSQL;
- команда: запуск приложения, миграции, SQL-запросы;
- окружение: PostgreSQL, DSN, переменные окружения.

## Причина

Приложение не может подключиться к базе или выполнить запрос. Причина может быть в неработающем сервере PostgreSQL, неверном DSN, неправильном пользователе, отсутствующей базе, непримененной миграции или неверной схеме таблиц.

## Решение

```bash
psql -U <user> -d <db-name>
```

Что проверить:

- запущен ли PostgreSQL;
- существует ли база данных;
- существует ли пользователь;
- совпадает ли host, port, user, password и dbname;
- применены ли миграции;
- не отличается ли имя таблицы от SQL-запроса.

## Как не повторить

- Хранить пример DSN без реальных секретов.
- Проверять подключение отдельно от приложения.
- Применять миграции до запуска API.
- Разделять ошибки подключения и ошибки SQL-запросов.

## Связанная теория

- [[01_Theory/06_Databases/PostgreSQL]]
- [[01_Theory/06_Databases/Connection_String]]
- [[01_Theory/06_Databases/Migrations]]
