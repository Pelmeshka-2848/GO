---
type: theory
area: Go
title: Go Toolchain
tags:
  - go
  - golang
  - toolchain
  - sdk
  - goroot
  - gopath
  - standard-library
---

# Go Toolchain

## Кратко

**Go toolchain** — это набор инструментов, который нужен для разработки, запуска, тестирования и сборки программ на Go.

В состав Go toolchain входят:

- команда `go`;
- компилятор;
- стандартная библиотека;
- форматтер;
- инструменты тестирования;
- инструменты сборки;
- инструменты управления модулями;
- служебные утилиты, которые вызываются через `go`.

## Основная идея

В Go большая часть повседневной работы делается через одну команду:

```bash
go <command> [arguments]
```

Например:

```bash
go version
go env
go run .
go test ./...
go build
```

Команда `go` является центральной точкой управления проектом.

## Go SDK

В учебном контексте часто говорят “установить Go SDK”. Обычно под этим понимают официальную установку Go, которая включает:

```text
Go distribution
├── go command
├── compiler
├── standard library
├── tools
└── documentation/runtime files
```

На Windows стандартный путь установки обычно выглядит так:

```text
C:\Program Files\Go
```

Внутри этой папки важны две директории:

```text
C:\Program Files\Go\bin
C:\Program Files\Go\src
```

Где:

- `bin` содержит исполняемые файлы, включая `go.exe`;
- `src` содержит исходники стандартной библиотеки Go.

## GOROOT

`GOROOT` — путь к установленной Go toolchain.

Пример:

```text
C:\Program Files\Go
```

Проверка:

```powershell
go env GOROOT
```

Важно:

- `GOROOT` должен указывать на корень установленного Go;
- не надо указывать `C:\Program Files\Go\bin`;
- IDE должна видеть не только `bin`, но и `src`.

Если GoLand не видит стандартные пакеты `runtime`, `testing`, `fmt`, значит часто проблема именно в неверном `GOROOT`.

## GOPATH

`GOPATH` — рабочая область Go, которая исторически использовалась для хранения исходников и зависимостей.

Сейчас в обычных проектах используется **Go Modules**, поэтому проект больше не обязан находиться внутри `GOPATH`.

Но `GOPATH` всё ещё нужен для:

- кэша;
- установленных CLI-бинарников;
- совместимости со старыми механизмами.

Проверка:

```powershell
go env GOPATH
```

На Windows часто используется путь вида:

```text
C:\Users\<user>\go
```

## GOMODCACHE

`GOMODCACHE` — директория, куда Go складывает скачанные зависимости.

Проверка:

```powershell
go env GOMODCACHE
```

Обычно путь находится внутри `GOPATH`:

```text
C:\Users\<user>\go\pkg\mod
```

## Стандартная библиотека

Стандартная библиотека — набор пакетов, которые поставляются вместе с Go.

Примеры стандартных пакетов:

| Пакет | Назначение |
|---|---|
| `fmt` | форматированный ввод/вывод |
| `runtime` | информация о рантайме Go |
| `testing` | тестирование |
| `os` | работа с ОС |
| `net/http` | HTTP-клиент и HTTP-сервер |
| `context` | контекст выполнения и отмена операций |
| `errors` | работа с ошибками |
| `encoding/json` | JSON-кодирование и декодирование |

Пример:

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.GOOS)
}
```

## runtime

`runtime` — стандартный пакет, который даёт доступ к информации о среде выполнения Go.

Примеры:

```go
runtime.GOOS
runtime.GOARCH
runtime.NumCPU()
```

`runtime.GOOS` показывает целевую операционную систему:

```text
windows
linux
darwin
```

`runtime.GOARCH` показывает архитектуру:

```text
amd64
arm64
386
```

## testing

`testing` — стандартный пакет для тестов.

Тестовые файлы имеют суффикс:

```text
_test.go
```

Тестовая функция должна начинаться с `Test`:

```go
func TestSomething(t *testing.T) {
    if 2+2 != 4 {
        t.Fatal("unexpected result")
    }
}
```

Запуск:

```bash
go test ./...
```

## Проверка установки Go

Минимальная диагностика:

```powershell
go version
where.exe go
go env GOROOT GOPATH GOMODCACHE GOPROXY GOOS GOARCH
```

Проверка стандартной библиотеки:

```powershell
$goRoot = go env GOROOT
dir "$goRoot\src\runtime"
dir "$goRoot\src\testing"
```

## Типовые проблемы

### go не найден

Симптом:

```text
go : The term 'go' is not recognized
```

Причины:

- Go не установлен;
- `C:\Program Files\Go\bin` не добавлен в `PATH`;
- терминал был открыт до установки Go.

Решение:

```powershell
go version
where.exe go
```

Если `where.exe go` ничего не показывает, нужно проверить `PATH`.

### IDE не видит runtime/testing

Симптомы:

```text
Cannot resolve symbol 'runtime'
Cannot resolve symbol 'testing'
Unresolved type 'T'
```

Причина:

- GoLand не видит SDK;
- неверно указан `GOROOT`;
- проект открыт не как Go-модуль;
- сломана индексация IDE.

Решение:

```text
File → Settings → Go → GOROOT
```

Указать:

```text
C:\Program Files\Go
```

После этого:

```text
File → Invalidate Caches → Invalidate and Restart
```

## Минимальный workflow

```text
установить Go
→ проверить go version
→ проверить go env
→ создать проект
→ go mod init
→ написать main.go
→ go run .
→ go test ./...
→ go build
```

## Команды для запоминания

```bash
go version
go env
go env GOROOT
go env GOPATH
go env GOMOD
go run .
go test ./...
go build
```

## Связанные заметки

- [[01_Theory/Go/001_Go_Modules]]
- [[01_Theory/Go/003_Go_CLI_Commands]]
- [[01_Theory/Go/004_GoLand_Setup]]

## Источники

- https://go.dev/doc/install
- https://go.dev/doc/toolchain
- https://pkg.go.dev/cmd/go
- https://pkg.go.dev/runtime
- https://pkg.go.dev/testing
