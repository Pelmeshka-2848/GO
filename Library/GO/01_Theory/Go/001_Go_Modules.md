---
type: theory
area: Go
title: Go Modules
tags:
  - go
  - golang
  - modules
  - go-mod
  - dependencies
  - gomod
---

# Go Modules

## Кратко

**Go Modules** — это современная система управления проектами и зависимостями в Go.

Модуль описывается файлом:

```text
go.mod
```

Обычно `go.mod` лежит в корне проекта.

Пример:

```text
hello-go/
├── go.mod
├── main.go
└── main_test.go
```

## Зачем нужны модули

Go Modules решают несколько задач:

- задают имя проекта через module path;
- фиксируют версию Go;
- хранят список зависимостей;
- позволяют хранить проект вне `GOPATH`;
- делают сборку проекта воспроизводимой;
- позволяют работать с версиями библиотек.

## Создание модуля

Команда:

```bash
go mod init example.com/sheve/hello-go
```

После выполнения создаётся файл:

```text
go.mod
```

Пример содержимого:

```go
module example.com/sheve/hello-go

go 1.xx.x
```

## module path

`module path` — имя модуля.

Пример:

```go
module example.com/sheve/hello-go
```

Это не обязательно должен быть реально существующий домен для учебной локальной лабораторной. Но в реальных проектах обычно используют путь, связанный с репозиторием:

```text
github.com/username/project
```

Пример:

```go
module github.com/stulevtoday/hello-go
```

## Структура go.mod

Простой файл:

```go
module example.com/sheve/hello-go

go 1.22
```

Файл с зависимостями:

```go
module example.com/sheve/api

go 1.22

require (
    github.com/gin-gonic/gin v1.10.0
    github.com/jackc/pgx/v5 v5.5.5
)
```

## Основные директивы go.mod

| Директива | Назначение |
|---|---|
| `module` | имя модуля |
| `go` | версия Go, на которую ориентируется модуль |
| `require` | зависимости модуля |
| `replace` | замена пути зависимости, часто для локальной разработки |
| `exclude` | исключение конкретной версии зависимости |
| `retract` | отзыв версии модуля автором |

## go.sum

При добавлении внешних зависимостей рядом с `go.mod` появляется файл:

```text
go.sum
```

`go.sum` хранит контрольные суммы зависимостей.

Важно:

- `go.sum` обычно коммитят в Git;
- файл помогает проверять целостность зависимостей;
- руками его обычно не редактируют.

## Добавление зависимости

В современных проектах чаще всего зависимость появляется автоматически после импорта и запуска команды:

```bash
go mod tidy
```

Также можно явно добавить зависимость:

```bash
go get github.com/gin-gonic/gin
```

## go mod tidy

Команда:

```bash
go mod tidy
```

Назначение:

- добавляет зависимости, которые реально используются в коде;
- удаляет зависимости, которые больше не используются;
- обновляет `go.sum`.

Типичный workflow:

```text
добавил import
→ go mod tidy
→ go test ./...
```

## go mod download

Команда:

```bash
go mod download
```

Назначение:

- заранее скачать зависимости;
- заполнить локальный кэш модулей.

Полезно в CI/CD или при подготовке окружения.

## go mod graph

Команда:

```bash
go mod graph
```

Показывает граф зависимостей.

Применяется для анализа:

- откуда пришла зависимость;
- какие есть транзитивные зависимости;
- почему проект тянет определённый пакет.

## go list -m all

Команда:

```bash
go list -m all
```

Показывает список всех модулей, участвующих в сборке.

## GOPROXY

`GOPROXY` — настройка источника, откуда Go скачивает модули.

Проверка:

```powershell
go env GOPROXY
```

Типичное значение:

```text
https://proxy.golang.org,direct
```

Установка:

```powershell
go env -w GOPROXY=https://proxy.golang.org,direct
```

## GOMOD

`GOMOD` показывает путь до текущего `go.mod`.

Команда:

```bash
go env GOMOD
```

Если ты находишься внутри Go-модуля, результат будет похож на:

```text
C:\GoLabs\000_Go_Environment_Setup\go.mod
```

Если модуля нет, Go покажет специальное пустое/служебное значение в зависимости от ОС и режима.

## Проект вне GOPATH

В современном Go нормальная структура такая:

```text
C:\GoLabs\my-project
├── go.mod
└── main.go
```

Проект **не обязан** лежать в:

```text
C:\Users\<user>\go\src
```

Это важное отличие Go Modules от старого GOPATH-подхода.

## Один модуль — один корень проекта

Обычно один репозиторий содержит один `go.mod` в корне:

```text
project/
├── go.mod
├── cmd/
├── internal/
└── pkg/
```

Но бывают multi-module repositories:

```text
repo/
├── service-a/
│   └── go.mod
└── service-b/
    └── go.mod
```

Для новичка лучше начинать с одного модуля на проект.

## Локальные replace

`replace` позволяет временно заменить зависимость локальным путём.

Пример:

```go
replace example.com/sheve/common => ../common
```

Используется, когда несколько локальных модулей разрабатываются рядом.

Для учебных первых лабораторных `replace` обычно не нужен.

## Минимальный workflow модуля

```bash
mkdir hello-go
cd hello-go
go mod init example.com/sheve/hello-go
go run .
go test ./...
go mod tidy
```

## Типовые ошибки

### go.mod file not found

Симптом:

```text
go: go.mod file not found in current directory or any parent directory
```

Причина:

- команда выполнена вне проекта;
- модуль не был инициализирован.

Решение:

```bash
go mod init example.com/sheve/hello-go
```

или перейти в папку, где уже лежит `go.mod`.

### package is not in std

Симптом:

```text
package example.com/... is not in std
```

Частая причина:

- неверный import path;
- модуль не инициализирован;
- проект открыт не из корня;
- IDE не видит `go.mod`.

## Что запомнить

- `go.mod` — главный файл Go-модуля.
- `go mod init` создаёт модуль.
- `go mod tidy` приводит зависимости в порядок.
- Проект не обязан лежать внутри `GOPATH`.
- `go.sum` нужен для контроля зависимостей.
- В реальных проектах module path часто совпадает с URL репозитория.

## Связанные заметки

- [[01_Theory/Go/000_Go_Toolchain]]
- [[01_Theory/Go/002_Go_Project_Structure]]
- [[01_Theory/Go/003_Go_CLI_Commands]]
- [[01_Theory/Go/004_GoLand_Setup]]

## Источники

- https://go.dev/ref/mod
- https://go.dev/doc/modules/gomod-ref
- https://go.dev/doc/modules/managing-dependencies
- https://go.dev/doc/tutorial/create-module
- https://go.dev/doc/code
