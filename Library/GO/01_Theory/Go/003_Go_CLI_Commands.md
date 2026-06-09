---
type: theory
area: Go
title: Go CLI Commands
tags:
  - go
  - golang
  - cli
  - commands
  - go-run
  - go-test
  - go-build
  - go-fmt
---

# Go CLI Commands

## Кратко

`go` — основная CLI-команда для работы с Go-проектами.

Общий формат:

```bash
go <command> [arguments]
```

Примеры:

```bash
go version
go env
go run .
go test ./...
go build
```

## go version

Команда:

```bash
go version
```

Показывает установленную версию Go.

Пример вывода:

```text
go version go1.xx.x windows/amd64
```

Используется для первой проверки установки.

## go env

Команда:

```bash
go env
```

Показывает переменные окружения Go.

Часто удобнее смотреть конкретные параметры:

```powershell
go env GOROOT GOPATH GOMODCACHE GOPROXY GOOS GOARCH
```

Полезные параметры:

| Параметр | Назначение |
|---|---|
| `GOROOT` | путь к Go SDK |
| `GOPATH` | рабочая область Go |
| `GOMOD` | путь к текущему `go.mod` |
| `GOMODCACHE` | кэш модулей |
| `GOPROXY` | источник загрузки модулей |
| `GOOS` | целевая ОС |
| `GOARCH` | целевая архитектура |

## go mod init

Команда:

```bash
go mod init example.com/sheve/hello-go
```

Создаёт новый Go-модуль и файл:

```text
go.mod
```

Используется один раз при создании проекта.

## go run

Команда:

```bash
go run .
```

Компилирует и запускает программу без явного сохранения бинарника в проекте.

Для простого проекта:

```bash
go run .
```

Для проекта с `cmd`:

```bash
go run ./cmd/api
```

Важно:

- `go run` всё равно компилирует код;
- удобно для разработки;
- для финальной сборки используют `go build`.

## go build

Команда:

```bash
go build
```

Компилирует проект.

С указанием имени файла:

```powershell
go build -o bin\hello-go.exe .
```

Для структуры `cmd`:

```powershell
go build -o bin\api.exe ./cmd/api
```

Отличие от `go run`:

| Команда | Назначение |
|---|---|
| `go run` | быстро запустить |
| `go build` | собрать бинарный файл |

## go test

Команда:

```bash
go test ./...
```

Запускает тесты во всех пакетах текущего модуля.

Разбор:

| Часть | Значение |
|---|---|
| `go test` | запустить тестирование |
| `./...` | все пакеты рекурсивно от текущей директории |

Тесты ищутся в файлах:

```text
*_test.go
```

Пример теста:

```go
package main

import "testing"

func TestSmoke(t *testing.T) {
}
```

## go fmt

Команда:

```bash
go fmt ./...
```

Форматирует Go-код стандартным форматтером.

В Go форматирование является частью культуры языка: обычно не спорят о стиле, а запускают стандартный форматтер.

Для одного файла:

```bash
go fmt main.go
```

Для всего проекта:

```bash
go fmt ./...
```

## gofmt

`gofmt` — отдельный форматтер.

Команда:

```bash
gofmt -w main.go
```

Разница:

| Команда | Назначение |
|---|---|
| `go fmt` | запускает форматирование на уровне пакетов |
| `gofmt` | форматирует конкретные файлы |

Для учебных лабораторных обычно достаточно:

```bash
go fmt ./...
```

## go mod tidy

Команда:

```bash
go mod tidy
```

Приводит зависимости в порядок:

- добавляет нужные;
- удаляет лишние;
- обновляет `go.sum`.

Типичный момент запуска:

```text
после добавления/удаления import
```

## go get

Команда:

```bash
go get github.com/gin-gonic/gin
```

Используется для добавления или изменения зависимости в текущем модуле.

Важно:

- для установки CLI-инструментов сейчас обычно используют `go install package@version`;
- для зависимостей проекта можно использовать `go get`.

## go install

Команда:

```bash
go install example.com/tool/cmd/tool@latest
```

Часто используется для установки CLI-инструментов.

Бинарник обычно попадает в:

```text
%GOPATH%\bin
```

Проверка:

```powershell
go env GOPATH
```

## go clean

Команда:

```bash
go clean
```

Удаляет временные файлы сборки.

Очистка кэша тестов:

```bash
go clean -testcache
```

После этого тесты будут выполнены заново без использования кэша.

## go list

Команда:

```bash
go list ./...
```

Показывает пакеты текущего модуля.

Для списка модулей:

```bash
go list -m all
```

## go doc

Команда:

```bash
go doc fmt.Println
```

Показывает документацию по пакету, типу или функции.

Примеры:

```bash
go doc fmt
go doc runtime.GOOS
go doc testing.T
```

## go help

Команда:

```bash
go help
```

Справка по командам.

Справка по конкретной команде:

```bash
go help build
go help test
go help mod
```

## Частые шаблоны команд

### Создать проект

```powershell
mkdir hello-go
cd hello-go
go mod init example.com/sheve/hello-go
```

### Запустить проект

```powershell
go run .
```

### Проверить всё

```powershell
go fmt ./...
go test ./...
go build
```

### Собрать .exe

```powershell
mkdir bin
go build -o bin\app.exe .
```

### Запустить собранный .exe

```powershell
.\bin\app.exe
```

### Проверить модуль

```powershell
go env GOMOD
```

### Проверить путь Go SDK

```powershell
go env GOROOT
```

## Разница между `.`, `./...` и путём

| Аргумент | Значение |
|---|---|
| `.` | текущий пакет |
| `./...` | все пакеты рекурсивно |
| `./cmd/api` | конкретный пакет в директории |
| `main.go` | конкретный файл |

Примеры:

```bash
go run .
go test ./...
go build ./cmd/api
```

## Типовой цикл разработки

```text
изменил код
→ go fmt ./...
→ go test ./...
→ go run .
→ go build
```

## Типовые ошибки

### go: go.mod file not found

Причина:

- команда запущена вне Go-модуля.

Проверка:

```bash
go env GOMOD
```

Решение:

```bash
go mod init example.com/sheve/project
```

или перейти в корень проекта.

### no Go files

Симптом:

```text
no Go files in ...
```

Причина:

- команда запущена в папке без `.go`-файлов;
- для проекта с `cmd` нужно указать правильный путь.

Решение:

```bash
go run ./cmd/api
```

### no test files

Симптом:

```text
[no test files]
```

Причина:

- в пакете нет файлов `*_test.go`.

Это не всегда ошибка. Для пакета без тестов такой вывод нормален.

## Что запомнить

- `go version` — проверить установку.
- `go env` — проверить окружение.
- `go mod init` — создать модуль.
- `go run .` — запустить.
- `go fmt ./...` — отформатировать.
- `go test ./...` — проверить тесты.
- `go build` — собрать.
- `go mod tidy` — привести зависимости в порядок.

## Связанные заметки

- [[01_Theory/Go/000_Go_Toolchain]]
- [[01_Theory/Go/001_Go_Modules]]
- [[01_Theory/Go/002_Go_Project_Structure]]

## Источники

- https://pkg.go.dev/cmd/go
- https://go.dev/doc/cmd
- https://go.dev/doc/tutorial/compile-install
- https://go.dev/doc/code
- https://go.dev/doc/modules/managing-dependencies
- https://go.dev/doc/go-get-install-deprecation
