---
type: rule
area: theory
status: active
created: 2026-06-06
updated: 2026-06-09
tags:
  - area/theory
  - type/rule
---

# Project Layout Rules

## Правило

Структура проекта должна расти из реальной сложности, а не копироваться заранее из больших production-шаблонов.

## Почему так

Лишние папки мешают учиться: становится непонятно, где находится логика и зачем нужен каждый слой. Маленький проект должен оставаться маленьким.

## Хорошо

```text
project/
├── go.mod
└── main.go
```

Для проекта с несколькими слоями:

```text
project/
├── cmd/app/main.go
├── internal/handler/
├── internal/service/
├── internal/repository/
└── go.mod
```

## Плохо

```text
project/
├── application/
├── domain/
├── infrastructure/
├── interfaces/
├── pkg/
├── platform/
└── main.go
```

Плохо не потому, что эти слова запрещены, а потому что структура появилась раньше реальной необходимости.

## Когда можно нарушить

- Если лабораторная специально про Clean Architecture.
- Если проект уже достаточно большой.
- Если структура задана учебным заданием.

## Связанные темы

- [[01_Theory/02_Go_Tooling/Project_Structure]]
- [[01_Theory/08_Architecture/Layered Architecture]]
- [[01_Theory/08_Architecture/Clean Architecture]]

