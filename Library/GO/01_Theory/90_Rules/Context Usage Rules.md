---
type: rule
area: theory
status: active
created: 2026-06-06
updated: 2026-06-09
tags:
  - area/theory
  - type/rule
---

# Context Usage Rules

## Правило

`context.Context` передается первым аргументом в функции, которые выполняют I/O, долгую работу или могут быть отменены.

## Почему так

Context дает единый способ отменить операцию, ограничить время выполнения и передать request-scoped значения через слои.

## Хорошо

```go
func (s *Service) GetUser(ctx context.Context, id int) (User, error) {
    return s.repo.FindUser(ctx, id)
}
```

## Плохо

```go
type Service struct {
    ctx context.Context
}
```

Хранить context в структуре обычно плохо: context относится к конкретной операции, а не к долгоживущему объекту.

## Ключевые правила

- Context первым аргументом.
- Не передавать `nil`.
- Не хранить context в struct без серьезной причины.
- После `context.WithCancel` или `context.WithTimeout` вызывать `cancel`.
- Не использовать context как обычную map для параметров.

## Когда можно нарушить

- В специальных инфраструктурных объектах, где жизненный цикл context действительно совпадает с объектом.
- В тестовых helper-ах, если это явно упрощает код и не попадает в production.

## Связанные темы

- [[01_Theory/07_Concurrency/Context]]
- [[01_Theory/07_Concurrency/Timeouts]]
- [[01_Theory/04_Stdlib/context package]]

