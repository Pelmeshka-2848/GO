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

# Pagination

## Кратко

Pagination — разбиение большого списка данных на страницы или порции.

## Зачем нужно

API не должен возвращать тысячи записей сразу: это нагружает БД, сеть, память и клиента.

## Основные подходы

- limit/offset;
- cursor pagination;
- keyset pagination.

Для первых лабораторных достаточно `limit` и `offset`.

## Правила

- Ограничивай максимальный `limit`.
- Задавай значение по умолчанию.
- Сортировка должна быть стабильной.
- Для больших таблиц offset может быть дорогим.
- В ответе полезно возвращать metadata: limit, offset, total или next cursor.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| Нет limit | Вернули слишком много | Ввести default/max limit |
| Нестабильный порядок | Нет ORDER BY | Добавить сортировку |
| Огромный offset | Большая таблица | Рассмотреть cursor/keyset |

## Связанные темы

- [[01_Theory/06_Databases/CRUD_SQL]]
- [[01_Theory/05_HTTP_Backend/02_API_Practices/REST API]]

