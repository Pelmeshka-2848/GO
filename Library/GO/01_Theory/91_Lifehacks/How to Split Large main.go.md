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

# How to Split Large main.go

## Суть лайфхака

Когда `main.go` становится длинным, выноси не “папки ради папок”, а конкретные ответственности.

## Когда использовать

- `main.go` содержит HTTP handlers;
- там же SQL;
- там же config parsing;
- там же бизнес-логика;
- файл трудно читать сверху вниз.

## Прием

1. Оставь в `main` запуск и сборку зависимостей.
2. Вынеси handlers в отдельный файл или пакет.
3. Вынеси бизнес-сценарии в service.
4. Вынеси SQL в repository.
5. Вынеси config parsing в отдельную функцию или пакет.

## Ограничения

- Для маленькой первой лабораторной один `main.go` нормален.
- Не создавай `internal/service` до появления реальной service-логики.

## Связанная теория

- [[01_Theory/02_Go_Tooling/Project_Structure]]
- [[01_Theory/08_Architecture/Layered Architecture]]
- [[01_Theory/90_Rules/Project Layout Rules]]

