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

# Logging Rules

## Правило

Лог должен помогать понять, что произошло, где произошло и с каким контекстом, но не должен раскрывать секреты.

## Почему так

Логи нужны для отладки и эксплуатации. Слишком мало контекста делает лог бесполезным, а лишние данные создают риск утечки.

## Хорошо

```text
level=error msg="create user failed" request_id=abc123 reason="email already exists"
```

## Плохо

```text
error
password=secret123 token=real-token
```

## Ключевые правила

- Не логируй пароли, токены и полные connection strings.
- Добавляй request ID или trace ID, если приложение обрабатывает HTTP-запросы.
- Логируй ошибку на границе приложения, а не на каждом уровне.
- Не превращай обычный control flow в поток error-логов.
- Для production лучше структурированные логи.

## Когда можно нарушить

- В локальной одноразовой отладке можно временно вывести больше данных, но такие логи не должны попадать в Git.

## Связанные темы

- [[01_Theory/09_Production/Logging]]
- [[01_Theory/09_Production/Secrets]]
- [[01_Theory/90_Rules/Error Handling Rules]]

