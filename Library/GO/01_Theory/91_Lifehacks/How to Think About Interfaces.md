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

# How to Think About Interfaces

## Суть лайфхака

Думай об interface как о потребности кода, а не как об обязательной паре к каждой struct.

## Когда использовать

- хочешь протестировать service без реальной БД;
- нужно заменить внешний API fake-клиентом;
- слой должен зависеть от поведения, а не от реализации.

## Прием

1. Сначала напиши конкретную реализацию.
2. Когда появится потребитель, посмотри, какие методы ему реально нужны.
3. Объяви маленький interface рядом с потребителем.
4. Не добавляй методы “на будущее”.

## Ограничения

- Если в проекте уже есть устойчивый публичный контракт, interface может жить отдельно.
- Для generated code и frameworks правила могут отличаться.

## Связанная теория

- [[01_Theory/01_Go_Basics/Interfaces]]
- [[01_Theory/90_Rules/Interface Design Rules]]
- [[01_Theory/08_Architecture/Dependency Injection]]

