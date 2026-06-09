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

# Hexagonal Architecture

## Кратко

Hexagonal Architecture, или Ports and Adapters, отделяет ядро приложения от внешних способов взаимодействия: HTTP, CLI, БД, очередей, внешних API.

## Зачем нужно

Подход помогает тестировать бизнес-логику без реальной инфраструктуры и менять внешние адаптеры без переписывания ядра.

## Основные понятия

- core — бизнес-логика;
- port — интерфейс, который описывает потребность core;
- adapter — реализация порта для HTTP, БД, API и т.д.

## Правила

- Core не зависит от adapter-ов.
- Adapter преобразует внешний формат в понятный core формат.
- Ports должны быть маленькими и осмысленными.
- Не нужно рисовать “шестиугольник” в коде буквально.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| Слишком абстрактные ports | Архитектура ради архитектуры | Делать interface по потребности |
| Core импортирует pgx/http | Зависимости направлены наружу | Ввести port |
| Адаптеры содержат business logic | Быстро написали | Перенести правило в core/service |

## Связанные темы

- [[01_Theory/08_Architecture/Clean Architecture]]
- [[01_Theory/08_Architecture/Dependency Injection]]
- [[01_Theory/90_Rules/Interface Design Rules]]

