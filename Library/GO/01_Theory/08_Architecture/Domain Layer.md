---
type: theory
area: theory
status: active
created: 2026-06-06
updated: 2026-06-09
tags:
  - area/theory
  - architecture
---

# Domain Layer

## Кратко

Domain layer содержит основные понятия предметной области: сущности, value objects, бизнес-правила и инварианты.

## Зачем нужно

Domain помогает не растворить бизнес-смысл в HTTP-handler-ах, SQL-запросах и DTO.

## Что может быть в domain

- сущности: User, Task, Order;
- статусы и enum-подобные типы;
- методы, которые меняют состояние сущности;
- ошибки предметной области;
- чистые бизнес-правила.

## Что не должно быть в domain

- `http.Request`;
- JSON-specific DTO;
- SQL rows;
- pgx/sql.DB;
- framework-specific код.

## Правила

- Domain должен быть максимально независимым.
- Не все struct автоматически являются domain model.
- Domain-методы должны защищать важные инварианты.
- Простые CRUD-проекты могут иметь тонкий domain.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| JSON tags везде | Одна struct для всего | Разделить DTO/domain |
| Domain без поведения | Только таблицы БД | Добавить методы там, где есть правила |
| HTTP ошибки в domain | Удобно вернуть статус | Вернуть доменную ошибку |

## Связанные темы

- [[01_Theory/08_Architecture/DTO]]
- [[01_Theory/08_Architecture/Service Layer]]
- [[01_Theory/08_Architecture/Clean Architecture]]

