# Lab 005 - PostgreSQL CRUD — Code

## Структура

```text
cmd/app           - точка входа приложения
internal/handler  - HTTP handlers
internal/repo     - работа с PostgreSQL
migrations        - SQL-миграции
```

## Основной код

Ключевой код этой лабораторной должен фиксироваться после самостоятельной реализации.

Что важно показать в разборе:

- где создается подключение к базе;
- где выполняются SQL-запросы;
- как HTTP-слой вызывает repository;
- как обрабатываются ошибки `not found`, `bad request` и ошибки базы.

## Команды запуска

```bash
go run .
```
