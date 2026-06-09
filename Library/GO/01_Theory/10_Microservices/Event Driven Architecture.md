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

# Event Driven Architecture

## Кратко

Event-driven architecture строит взаимодействие вокруг событий: что-то произошло, другие части системы реагируют.

## Зачем нужно

- слабее связывать сервисы;
- выполнять фоновые процессы;
- реагировать на изменения;
- масштабировать consumers независимо;
- строить audit/event history.

## Примеры событий

- `UserRegistered`;
- `OrderPaid`;
- `TaskCompleted`;
- `PasswordResetRequested`.

## Правила

- Событие должно описывать факт, а не команду.
- Контракт события должен быть стабильным.
- Consumer должен быть идемпотентным.
- Не все процессы нужно делать event-driven.
- Нужна observability обработки событий.

## Типичные ошибки

- использовать события как скрытые RPC-команды;
- не учитывать повторы;
- не иметь порядка версий события;
- усложнить простой synchronous flow.

## Связанные темы

- [[01_Theory/10_Microservices/Message Broker]]
- [[01_Theory/10_Microservices/Metrics and Observability]]

