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

# How to Design Small Packages

## Суть лайфхака

Package должен иметь понятную ответственность и короткое имя. Хороший package легко описать одной фразой.

## Когда использовать

- проект вырос из одного файла;
- появились повторяющиеся функции;
- хочется отделить domain, service, repository;
- код сложно тестировать.

## Прием

- Называй package по тому, что он делает: `user`, `task`, `config`, `storage`.
- Не называй package общими словами: `utils`, `common`, `helpers`, если можно конкретнее.
- Держи публичный API package маленьким.
- Не экспортируй то, что нужно только внутри package.

## Ограничения

- В учебном проекте можно начать проще.
- Один маленький helper не всегда заслуживает отдельный package.

## Связанная теория

- [[01_Theory/90_Rules/Go Naming Rules]]
- [[01_Theory/02_Go_Tooling/Project_Structure]]

