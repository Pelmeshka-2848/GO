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

# How to Work Without Docker

## Суть лайфхака

Если Docker пока мешает обучению, можно сначала запускать Go-приложение и PostgreSQL локально, но нужно аккуратно фиксировать окружение.

## Когда использовать

- изучаешь базовый Go;
- делаешь первые CLI/HTTP лабораторные;
- Docker создает больше проблем, чем пользы;
- нужно быстрее разобраться в коде.

## Прием

- Запиши версии Go и PostgreSQL.
- Храни настройки в `.env`, но не коммить его.
- Сделай `.env.example`.
- Фиксируй команды запуска в worklog.
- Не смешивай локальные пути и production assumptions.

## Ограничения

- Для микросервисов Docker или другой способ изоляции окружения быстро станет полезным.
- Для командной разработки локальные ручные настройки хуже воспроизводятся.

## Связанная теория

- [[01_Theory/09_Production/Config_Env]]
- [[01_Theory/06_Databases/PostgreSQL]]
- [[02_Practice/03_Debug_Diary/Windows Environment Problems]]

