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

# Clean Architecture

## Кратко

Clean Architecture — подход, где бизнес-правила не зависят от внешних деталей: HTTP, базы данных, framework-ов и UI.

## Зачем нужно

- защитить бизнес-логику от инфраструктурных изменений;
- упростить тестирование use cases;
- сделать зависимости направленными внутрь;
- отделить домен от delivery и persistence.

## Основная идея

Внутренние слои не знают о внешних. HTTP и база данных — детали, подключенные снаружи.

```text
transport -> usecase/service -> domain
          -> repository interface
infrastructure implements repository
```

## Правила

- Domain/usecase не импортирует HTTP.
- Domain/usecase не импортирует SQL-driver.
- Interfaces часто объявляются на стороне потребителя.
- DTO не должен протекать глубоко в domain без необходимости.
- Для маленьких проектов clean-подход можно применять частично.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| Слишком много папок | Копия шаблона | Начать с простых слоев |
| Domain знает JSON tags | API протек внутрь | Ввести DTO/mapping |
| Usecase зависит от pgx | Infrastructure протекла внутрь | Зависеть от interface |

## Когда не нужно

Для маленькой учебной CLI-программы полный Clean Architecture будет лишним. Достаточно функций и простых пакетов.

## Связанные темы

- [[01_Theory/08_Architecture/Layered Architecture]]
- [[01_Theory/08_Architecture/Dependency Injection]]
- [[01_Theory/90_Rules/Project Layout Rules]]

