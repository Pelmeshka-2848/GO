---
type: theory
area: theory
status: active
created: 2026-06-06
updated: 2026-06-09
tags:
  - area/theory
  - microservices
---

# What Is Microservice

## Кратко

Microservice — небольшой сервис, отвечающий за отдельную бизнес-возможность и разворачиваемый независимо.

## Зачем нужно

- независимые релизы;
- масштабирование отдельных частей;
- разделение ответственности между командами;
- изоляция технологических решений;
- устойчивость при правильном проектировании.

## Основная идея

Сервис должен иметь понятную границу ответственности. “User service” или “Billing service” обычно осмысленнее, чем “Database service” или “Handlers service”.

## Цена микросервисов

- сетевые ошибки;
- distributed tracing;
- eventual consistency;
- retries и timeouts;
- service discovery;
- больше инфраструктуры;
- сложнее локальная разработка.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| Разделить по слоям | HTTP отдельно, DB отдельно | Делить по бизнес-возможностям |
| Общая БД на всех | Удобно стартовать | Явные ownership-правила |
| Нет observability | “Потом добавим” | Логи/метрики/tracing с начала |

## Связанные темы

- [[01_Theory/08_Architecture/Microservices]]
- [[01_Theory/10_Microservices/Distributed Tracing]]

