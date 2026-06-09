---
type: theory
area: theory
status: active
created: 2026-06-06
updated: 2026-06-09
tags:
  - area/theory
  - backend
  - go
---

# Graceful Shutdown

## Кратко

Graceful shutdown — корректное завершение сервера: перестать принимать новые запросы, дать текущим операциям закончиться и освободить ресурсы.

## Зачем нужно

Без graceful shutdown сервер может оборвать запросы, потерять работу worker-ов или не закрыть подключения к БД.

## Что обычно закрывать

- HTTP server;
- database pool;
- background workers;
- message consumers;
- loggers/exporters.

## Основная идея

Приложение слушает системный сигнал, создает context с timeout и вызывает shutdown-методы компонентов.

## Правила

- Shutdown должен иметь timeout.
- Новые запросы больше не принимаются.
- Долгие операции должны уважать context.
- Ресурсы закрываются в понятном порядке.
- Ошибки shutdown логируются.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| Просто `Ctrl+C` | Dev-привычка | Обработать signal |
| Нет timeout | Shutdown может зависнуть | Context with timeout |
| Workers не остановлены | Нет контракта отмены | Передать context |

## Связанные темы

- [[01_Theory/09_Production/Graceful_Shutdown]]
- [[01_Theory/07_Concurrency/Context]]
- [[01_Theory/04_Stdlib/net-http package]]

