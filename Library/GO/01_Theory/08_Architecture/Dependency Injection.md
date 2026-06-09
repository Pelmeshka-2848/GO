---
type: theory
area: theory
status: active
created: 2026-06-06
updated: 2026-06-09
tags:
  - area/theory
  - architecture
---

# Dependency Injection

## Кратко

Dependency injection — передача зависимостей снаружи, а не создание их внутри объекта.

## Зачем нужно

- упростить тестирование;
- заменить реализацию без переписывания business logic;
- явно видеть зависимости компонента;
- не создавать глобальное состояние.

## Простая форма в Go

```go
type Service struct {
    repo UserRepository
}

func NewService(repo UserRepository) *Service {
    return &Service{repo: repo}
}
```

## Правила

- Зависимости передаются через конструктор или поля настройки.
- Не создавай DB/client внутри service, если service должен тестироваться.
- Interface нужен там, где есть реальная потребность в подмене поведения.
- Для маленького проекта достаточно ручного DI.
- Избегай глобальных singleton-ов без необходимости.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| Service сам открывает БД | Так проще начать | Передать repository |
| Interface для всего | “Так архитектурнее” | Делать interface у потребителя |
| Скрытые globals | Быстро удобно | Явные зависимости |

## Связанные темы

- [[01_Theory/90_Rules/Interface Design Rules]]
- [[01_Theory/08_Architecture/Service Layer]]
- [[01_Theory/01_Go_Basics/Interfaces]]

