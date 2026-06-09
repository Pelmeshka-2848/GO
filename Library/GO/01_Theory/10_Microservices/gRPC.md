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

# gRPC

## Кратко

gRPC — RPC-framework для межсервисного взаимодействия, обычно использующий Protocol Buffers как IDL и формат сериализации.

## Зачем нужно

- строгий контракт между сервисами;
- генерация client/server кода;
- эффективная бинарная сериализация;
- streaming;
- удобное internal API.

## Когда уместно

- внутренние сервисы;
- строгие контракты;
- высокая частота вызовов;
- streaming-сценарии;
- polyglot-среда.

## Когда REST проще

- публичный API для браузера;
- простая CRUD-лабораторная;
- ручное тестирование через curl;
- команда еще не готова к proto/tooling.

## Правила

- `.proto` файл является контрактом.
- Изменения контракта должны быть совместимыми.
- Нужны deadlines/timeouts.
- Ошибки должны маппиться в понятные gRPC status codes.
- Generated-код не редактируется вручную.

## Типичные ошибки

- использовать gRPC в маленьком проекте без пользы;
- ломать совместимость proto;
- не ставить deadlines;
- смешивать generated-код и ручную бизнес-логику.

## Связанные темы

- [[01_Theory/11_Advanced_Go/Code_Generation]]
- [[01_Theory/10_Microservices/Service Discovery]]


