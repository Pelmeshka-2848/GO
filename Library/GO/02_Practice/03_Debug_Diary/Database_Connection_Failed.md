# Ошибка подключения к базе данных

## Контекст

Возникает при запуске backend-приложения, когда оно не может подключиться к базе данных.

## Текст ошибки

```text
connection refused
database does not exist
password authentication failed
```

## Причина

База не запущена, неверно указан DSN, пользователь или пароль не совпадают, база не создана или миграции не применены.

## Решение

- Проверить, запущен ли PostgreSQL.
- Проверить host, port, user, password и dbname.
- Подключиться к базе отдельно через `psql`.
- Применить миграции перед запуском API.

## Как не повторить

- Хранить пример конфигурации без реальных секретов.
- Проверять подключение отдельно от HTTP API.
- Фиксировать команды подготовки базы в worklog.

## Связанная практика

- [[02_Practice/03_Debug_Diary/PostgreSQL Connection Errors]]
- [[02_Practice/04_Code_Recipes/Postgres_Connection]]
- [[01_Theory/06_Databases/Connection_String]]
