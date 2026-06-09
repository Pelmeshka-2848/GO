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

# Layered Architecture

## Кратко

Layered architecture делит приложение на слои с разными обязанностями: transport, service, repository, domain/model.

## Зачем нужно

- уменьшить смешение HTTP, бизнес-логики и SQL;
- упростить тестирование;
- сделать код понятнее;
- облегчить замену хранения или transport-слоя.

## Типовая схема

```text
handler -> service -> repository -> database
```

Handler принимает HTTP. Service выполняет сценарий. Repository работает с данными.

## Правила

- Верхний слой может вызывать нижний, но не наоборот.
- Handler не должен знать SQL.
- Repository не должен знать HTTP.
- Service не должен зависеть от конкретного router-а.
- Модели на границах можно разделять через DTO.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| SQL в handler | Быстрее написать | Вынести repository |
| Service возвращает HTTP status | Смешение слоев | Преобразовать ошибку в handler |
| Слишком много слоев | Копирование шаблона | Упростить структуру |

## Связанные темы

- [[01_Theory/08_Architecture/Service Layer]]
- [[01_Theory/06_Databases/Repository_Layer]]
- [[01_Theory/90_Rules/Project Layout Rules]]

