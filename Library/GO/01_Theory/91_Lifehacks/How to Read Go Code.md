---
type: lifehack
area: theory
status: active
created: 2026-06-06
updated: 2026-06-09
tags:
  - area/theory
  - type/lifehack
---

# How to Read Go Code

## Суть лайфхака

Читай Go-код от точки входа к границам: `main`, маршруты, handlers, services, repositories, внешние зависимости.

## Когда использовать

- когда открываешь новый проект;
- когда потерялся в структуре;
- когда нужно понять, как запрос проходит через приложение.

## Прием

1. Найди `go.mod` и имя модуля.
2. Найди `cmd/.../main.go` или корневой `main.go`.
3. Посмотри, где создаются зависимости.
4. Найди регистрацию routes/handlers.
5. Пройди один сценарий от handler-а до repository.
6. Отдельно выпиши внешние зависимости: БД, API, env.

## Ограничения

- В сильно event-driven системе точка входа может быть consumer, а не HTTP handler.
- В маленькой лабораторной вся логика может быть в одном файле, но порядок чтения все равно помогает.

## Связанная теория

- [[01_Theory/02_Go_Tooling/Project_Structure]]
- [[01_Theory/08_Architecture/Layered Architecture]]

