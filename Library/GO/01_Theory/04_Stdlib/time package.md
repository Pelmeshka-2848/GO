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

# time package

## Кратко

`time` — стандартный пакет для работы со временем, датами, duration, timers и tickers.

## Зачем нужно

- ставить timeout;
- измерять длительность операции;
- форматировать дату;
- планировать повторяющиеся действия;
- хранить timestamps.

## Основные типы

- `time.Time` — конкретный момент времени;
- `time.Duration` — длительность;
- `time.Timer` — одноразовый таймер;
- `time.Ticker` — повторяющийся тикер.

## Форматирование

Go использует особый layout:

```go
"2006-01-02 15:04:05"
```

Это не произвольные буквы, а референсная дата Go.

## Правила

- Для timeout в API чаще используй context.
- Для измерения времени используй `time.Since(start)`.
- Для формата даты используй Go layout, а не `YYYY-MM-DD`.
- Останавливай ticker, когда он больше не нужен.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| `YYYY-MM-DD` | Go layout работает иначе | Использовать `2006-01-02` |
| Не остановить ticker | Goroutine/resource leak | `defer ticker.Stop()` |
| Путать Time и Duration | Разные смыслы | Выбирать тип по задаче |

## Связанные темы

- [[01_Theory/99_Atomic_Go/011_Time_Package]]
- [[01_Theory/99_Atomic_Go/012_Time_Format_Layout]]
- [[01_Theory/07_Concurrency/Timeouts]]

