---
type: recipe
area: practice
status: draft
created: 2026-06-06
updated:
tags:
  - area/practice
  - type/snippet
---

# Recipe - Graceful Shutdown

## Когда использовать

- Когда HTTP-сервер должен завершаться без обрыва активных запросов.
- Когда приложение должно корректно реагировать на `Ctrl+C` или сигнал остановки.
- Когда есть ресурсы, которые нужно закрыть при завершении.

## Код

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

return server.Shutdown(ctx)
```

## Как подключить в проект

Обычно shutdown запускается после получения OS-сигнала. Сервер получает контекст с таймаутом и время на завершение текущих запросов.

## Связанная теория

- [[01_Theory/09_Production/Graceful_Shutdown]]
- [[01_Theory/07_Concurrency/Context]]

## Где применял

- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/011_Production_Ready_API/00_Lab Card]]
- [[02_Practice/04_Code_Recipes/Graceful_Shutdown]]
