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

# Go Compiler Errors

## Симптомы

```text
undefined: name
declared and not used
cannot use value as type
no required module provides package
```

## Где возникло

- проект: любой Go-проект или лабораторная;
- команда: `go run`, `go build`, `go test`;
- окружение: Go toolchain, Go modules, IDE.

## Причина

Компилятор Go строго проверяет код до запуска. Частые причины: неиспользуемые переменные, неверные типы, отсутствующие imports, неправильный module path, обращение к несуществующему имени.

## Решение

```bash
go fmt ./...
go test ./...
go mod tidy
```

Дополнительно:

- прочитать первую ошибку сверху вниз;
- проверить имя пакета и imports;
- проверить, совпадает ли тип аргумента с ожидаемой сигнатурой;
- убрать неиспользуемые переменные или временно заменить имя на `_`.

## Как не повторить

- Запускать `go test ./...` после небольших изменений.
- Не накапливать много правок перед первой компиляцией.
- Держать `go.mod` в корне проекта.

## Связанная теория

- [[01_Theory/02_Go_Tooling/Go_Build]]
- [[01_Theory/02_Go_Tooling/Go_Test]]
- [[01_Theory/02_Go_Tooling/Go_Mod_Tidy]]
