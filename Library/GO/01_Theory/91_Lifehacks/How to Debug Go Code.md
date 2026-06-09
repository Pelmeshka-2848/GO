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

# How to Debug Go Code

## Суть лайфхака

Сначала сузь место ошибки, потом проверяй гипотезы маленькими шагами. Не меняй сразу много кода.

## Когда использовать

- программа не компилируется;
- результат не совпадает с ожиданием;
- HTTP endpoint возвращает не тот статус;
- тест падает непонятно почему.

## Прием

1. Прочитай текст ошибки полностью.
2. Найди файл и строку.
3. Определи тип проблемы: compile error, runtime error, logic error, environment.
4. Воспроизведи минимальный сценарий.
5. Добавь временный вывод или тест.
6. После исправления запиши ошибку в debug diary.

## Ограничения

- `fmt.Println` нормален для первых шагов, но не должен заменять логи и тесты в backend-проекте.
- Если ошибка плавающая, проверь race, timing и внешние зависимости.

## Связанная теория

- [[01_Theory/90_Rules/Error Handling Rules]]
- [[01_Theory/07_Concurrency/Data_Race]]
- [[02_Practice/03_Debug_Diary/00_Debug_Diary_Index]]

