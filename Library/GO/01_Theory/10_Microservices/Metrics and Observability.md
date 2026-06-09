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

# Metrics and Observability

## Кратко

Metrics and observability в микросервисах помогают понимать состояние каждого сервиса и всей цепочки взаимодействий.

## Зачем нужно

Без метрик и наблюдаемости микросервисная система быстро становится черным ящиком.

## Что измерять

- request rate;
- error rate;
- latency;
- saturation ресурсов;
- queue length;
- consumer lag;
- dependency failures;
- retry count.

## Что связывать

- logs с trace id;
- metrics с service name;
- traces с endpoint-ами;
- alerts с runbook-ами.

## Правила

- У каждого сервиса должны быть базовые RED/USE metrics.
- Ошибки downstream-сервисов должны быть видимы.
- Alerts должны быть actionable.
- Dashboard не заменяет alerting.
- Нужны service labels и environment labels.

## Типичные ошибки

- логи есть, но нет trace id;
- метрики собираются, но нет alerts;
- невозможно понять, какой сервис виноват;
- не измеряется очередь или broker lag.

## Связанные темы

- [[01_Theory/09_Production/Observability]]
- [[01_Theory/10_Microservices/Distributed Tracing]]

