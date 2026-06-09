---
type: debug
area: practice
status: active
created: 2026-06-06
updated:
tags:
  - area/practice
  - type/error
---

# GoLand Problems

## Симптомы

```text
cannot resolve symbol
go.mod file not found
GOROOT is not defined
package is not in GOROOT
```

## Где возникло

- проект: учебные Go-проекты;
- команда: запуск из GoLand, индексирование, тесты из IDE;
- окружение: GoLand, Go SDK, Go modules.

## Причина

GoLand может не видеть SDK, открыть не корень проекта, использовать неправильный working directory или не успеть переиндексировать зависимости после изменения `go.mod`.

## Решение

```bash
go version
go mod tidy
go test ./...
```

Что проверить в IDE:

- выбран ли Go SDK;
- открыт ли корень проекта с `go.mod`;
- совпадает ли working directory с корнем проекта;
- подтянулись ли зависимости после `go mod tidy`.

## Как не повторить

- Открывать в GoLand папку проекта, а не случайную вложенную директорию.
- После изменения зависимостей запускать `go mod tidy`.
- Проверять проблему в терминале, чтобы отделить ошибку IDE от ошибки проекта.

## Связанная теория

- [[01_Theory/99_Atomic_Go/004_GoLand_Setup]]
- [[01_Theory/02_Go_Tooling/Go_Modules]]
- [[01_Theory/02_Go_Tooling/Project_Structure]]
