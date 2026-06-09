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

# Message Broker

## Кратко

Message broker — компонент для передачи сообщений между producer и consumer.

## Зачем нужно

Брокер помогает строить асинхронные процессы, разгружать request path и связывать сервисы через события.

## Что дает

- очередь задач;
- pub/sub;
- retry;
- буферизацию нагрузки;
- decoupling сервисов.

## Важные вопросы

- доставка at least once или at most once;
- порядок сообщений;
- повторная обработка;
- dead-letter queue;
- формат события;
- мониторинг lag/queue length.

## Правила

- Consumer должен быть готов к повторной доставке.
- События должны иметь версионируемый контракт.
- Ошибки обработки не должны исчезать молча.
- Нужна стратегия retry и DLQ.
- Брокер не заменяет нормальную модель данных.

## Типичные ошибки

- считать сообщение обработанным до успешной операции;
- не делать idempotency;
- менять event schema без совместимости;
- не мониторить очередь.

## Связанные темы

- [[01_Theory/08_Architecture/Message_Queues]]
- [[01_Theory/10_Microservices/Event Driven Architecture]]

