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

# Service Discovery

## Кратко

Service discovery — механизм, который помогает сервисам находить адреса друг друга в динамической инфраструктуре.

## Зачем нужно

В микросервисах адреса instances могут меняться: контейнеры перезапускаются, масштабируются, переезжают между узлами.

## Подходы

- DNS-based discovery;
- registry service;
- discovery через orchestrator;
- static config для простых учебных случаев.

## Правила

- Для production не полагайся на hardcoded IP.
- Discovery должен работать вместе с healthcheck.
- Client должен иметь timeout и retry policy.
- Нужна observability межсервисных вызовов.

## Типичные ошибки

- зашить адрес сервиса в код;
- не учитывать перезапуск instances;
- не проверять health сервиса;
- делать бесконечный retry на недоступный сервис.

## Связанные темы

- [[01_Theory/10_Microservices/API Gateway]]
- [[01_Theory/09_Production/Healthcheck]]

