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

# Distributed Tracing

## Кратко

Distributed tracing показывает путь запроса через несколько сервисов.

## Зачем нужно

В микросервисах один пользовательский запрос может пройти через gateway, auth, user service, billing service и broker. Без tracing трудно понять, где задержка или ошибка.

## Основные понятия

- trace — весь путь запроса;
- span — отдельный участок работы;
- trace id — общий идентификатор;
- span id — идентификатор участка;
- propagation — передача контекста между сервисами.

## Правила

- Trace context должен передаваться через межсервисные вызовы.
- Важные операции должны создавать spans.
- Errors должны отмечаться в trace.
- Tracing не заменяет logs и metrics.
- Не добавляй чувствительные данные в span attributes.

## Типичные ошибки

- не пробрасывать trace id дальше;
- иметь traces без полезных attributes;
- логировать request id отдельно от trace id без связи;
- собирать слишком много данных и перегрузить систему.

## Связанные темы

- [[01_Theory/10_Microservices/Metrics and Observability]]
- [[01_Theory/09_Production/Observability]]

