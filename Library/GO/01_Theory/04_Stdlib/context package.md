---
type: theory
area: theory
status: active
created: 2026-06-06
updated: 2026-06-09
tags:
  - area/theory
  - go
  - stdlib
---

# context package

## Кратко

`context` — стандартный пакет для передачи отмены, deadline, timeout и request-scoped values.

## Зачем нужно

- отменять долгие операции;
- ограничивать время выполнения;
- пробрасывать отмену HTTP-запроса в БД или внешний API;
- корректно завершать goroutines.

## Основные функции

- `context.Background`;
- `context.TODO`;
- `context.WithCancel`;
- `context.WithTimeout`;
- `context.WithDeadline`;
- `context.WithValue`.

## Главный паттерн

```go
ctx, cancel := context.WithTimeout(parent, timeout)
defer cancel()
```

`cancel` освобождает ресурсы, связанные с context, поэтому его нужно вызвать.

## Правила

- Context передается первым аргументом.
- Context не хранится в структуре без сильной причины.
- Не передавай `nil` context.
- `WithValue` не используется для обычных параметров функции.
- I/O-операции должны принимать context, если могут зависнуть.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| Не вызвать `cancel` | Ресурсы живут дольше нужного | `defer cancel()` |
| Создать `Background` внутри слоя | Теряется отмена родителя | Передавать существующий ctx |
| Хранить ctx в struct | Context относится к операции | Передавать аргументом |

## Связанные темы

- [[01_Theory/07_Concurrency/Context]]
- [[01_Theory/90_Rules/Context Usage Rules]]
- [[01_Theory/07_Concurrency/Timeouts]]

