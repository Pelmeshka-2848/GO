---
type: theory
area: Go
title: GoLand Setup
tags:
  - go
  - goland
  - jetbrains
  - ide
  - goroot
  - modules
  - sdk
---

# GoLand Setup

## Кратко

GoLand — IDE от JetBrains для разработки на Go.

Для корректной работы GoLand должен видеть:

- Go SDK;
- `GOROOT`;
- файл `go.mod`;
- Go Modules;
- стандартную библиотеку;
- структуру проекта.

Если IDE не видит `runtime`, `testing`, `fmt` и другие стандартные пакеты, проблема обычно не в коде, а в настройке SDK или индексации.

## Главная настройка: GOROOT

`GOROOT` — путь к установленному Go SDK.

В GoLand открыть:

```text
File → Settings → Go → GOROOT
```

На Windows обычно нужно указать:

```text
C:\Program Files\Go
```

Важно:

```text
правильно:   C:\Program Files\Go
неправильно: C:\Program Files\Go\bin
```

GoLand должен видеть внутри SDK папки:

```text
bin
src
pkg
```

## Как добавить SDK

Путь:

```text
File → Settings → Go → GOROOT
```

Дальше:

```text
Add SDK → Local
```

Выбрать:

```text
C:\Program Files\Go
```

Если Go не установлен, можно использовать вариант:

```text
Add SDK → Download
```

Но для стабильного учебного окружения лучше понимать, где физически установлен SDK.

## Проверка SDK через PowerShell

Перед настройкой IDE полезно проверить Go вне GoLand:

```powershell
go version
where.exe go
go env GOROOT GOPATH GOMODCACHE GOPROXY GOOS GOARCH
```

Если эти команды не работают, GoLand тоже не сможет нормально работать с проектом.

## Go Modules в GoLand

Для обычных современных проектов нужен режим Go Modules.

Путь:

```text
File → Settings → Go → Go Modules
```

Должна быть включена опция:

```text
Enable Go modules integration
```

Проект должен быть открыт из корня, где лежит:

```text
go.mod
```

Правильно:

```text
Open → C:\GoLabs\000_Go_Environment_Setup
```

Неправильно:

```text
Open → C:\GoLabs\000_Go_Environment_Setup\main.go
```

## Структура проекта для GoLand

Минимальный проект:

```text
000_Go_Environment_Setup/
├── go.mod
├── main.go
└── main_test.go
```

GoLand должен распознать:

- `go.mod` как модуль;
- `main.go` как исполняемый пакет;
- `main_test.go` как тестовый файл;
- стандартные пакеты из `GOROOT/src`.

## Run Configuration

Для запуска `main.go` можно использовать:

```text
Run → Run...
```

или кнопку запуска рядом с:

```go
func main()
```

Тип конфигурации:

```text
Go Build
```

Для простого проекта рабочий пакет:

```text
.
```

или корневая папка модуля.

Для структуры с `cmd`:

```text
./cmd/api
```

## Test Configuration

Для тестов можно использовать:

```text
Run → Run...
```

или кнопку запуска рядом с тестовой функцией:

```go
func TestSomething(t *testing.T)
```

CLI-эквивалент:

```powershell
go test ./...
```

## GOPATH в GoLand

В современных проектах на Go Modules не нужно помещать проект внутрь GOPATH.

Но GoLand всё равно может показывать настройки GOPATH.

Важно:

- `GOROOT` — это SDK;
- `GOPATH` — это рабочая область/кэш;
- `go.mod` — это корень современного проекта.

Не надо путать:

```text
GOROOT ≠ GOPATH
```

## Типовая проблема: Cannot resolve symbol runtime/testing

Симптомы:

```text
Cannot resolve symbol 'runtime'
Cannot resolve symbol 'testing'
Unresolved type 'T'
Unresolved reference 'GOOS'
Unresolved reference 'Fatal'
Unresolved reference 'GOARCH'
```

Что это значит:

- GoLand не видит стандартную библиотеку;
- SDK не настроен;
- выбран неправильный путь к SDK;
- индекс IDE повреждён;
- проект открыт не как Go-модуль.

Проверка через терминал:

```powershell
go test ./...
```

Если терминал работает, а GoLand подсвечивает ошибки — проблема в IDE.

## Решение проблемы runtime/testing

1. Проверить Go:

```powershell
go version
where.exe go
go env GOROOT
```

2. Проверить наличие стандартной библиотеки:

```powershell
$goRoot = go env GOROOT
dir "$goRoot\src\runtime"
dir "$goRoot\src\testing"
```

3. В GoLand открыть:

```text
File → Settings → Go → GOROOT
```

4. Указать:

```text
C:\Program Files\Go
```

5. Проверить Go Modules:

```text
File → Settings → Go → Go Modules
```

6. При необходимости сбросить индексы:

```text
File → Invalidate Caches → Invalidate and Restart
```

## Настройка GOPROXY в GoLand

Если есть проблемы с загрузкой зависимостей, можно указать переменную окружения:

```text
GOPROXY=https://proxy.golang.org,direct
```

Но сначала лучше проверить через PowerShell:

```powershell
go env GOPROXY
```

Установить глобально:

```powershell
go env -w GOPROXY=https://proxy.golang.org,direct
```

## Автоматическое форматирование

GoLand умеет форматировать код.

Но полезно привыкнуть к CLI-команде:

```powershell
go fmt ./...
```

Перед коммитом удобно выполнять:

```powershell
go fmt ./...
go test ./...
```

## Кэш и индексация

Если GoLand показывает ошибки, которых нет в терминале:

```text
File → Invalidate Caches → Invalidate and Restart
```

Это помогает при:

- неверной подсветке;
- зависших unresolved symbols;
- неправильном состоянии индекса;
- обновлении SDK;
- переносе проекта.

## Чек-лист корректной настройки

| Проверка | Как проверить |
|---|---|
| Go установлен | `go version` |
| go.exe найден | `where.exe go` |
| GOROOT корректный | `go env GOROOT` |
| Стандартная библиотека есть | `dir "$goRoot\src\runtime"` |
| GoLand видит SDK | `Settings → Go → GOROOT` |
| Проект открыт из корня | в корне есть `go.mod` |
| Go Modules включены | `Settings → Go → Go Modules` |
| Тесты работают | `go test ./...` |

## Рекомендуемый порядок диагностики

```text
1. Проверить go version в PowerShell
2. Проверить where.exe go
3. Проверить go env GOROOT
4. Проверить папки runtime/testing в GOROOT/src
5. Проверить GOROOT в GoLand
6. Проверить Go Modules
7. Открыть проект из корня
8. Invalidate Caches
9. Запустить go test ./...
```

## Что запомнить

- GoLand не заменяет Go SDK: SDK должен быть установлен или скачан.
- `GOROOT` указывает на корень SDK, не на `bin`.
- Проект надо открывать из папки с `go.mod`.
- `runtime` и `testing` — стандартные пакеты, они не устанавливаются отдельно.
- Если CLI работает, а IDE нет — почти всегда проблема в настройке GoLand или индексации.

## Связанные заметки

- [[01_Theory/Go/000_Go_Toolchain]]
- [[01_Theory/Go/001_Go_Modules]]
- [[01_Theory/Go/002_Go_Project_Structure]]
- [[01_Theory/Go/003_Go_CLI_Commands]]

## Источники

- https://www.jetbrains.com/help/go/configuring-goroot-and-gopath.html
- https://www.jetbrains.com/help/go/create-a-project-with-go-modules-integration.html
- https://www.jetbrains.com/help/go/getting-started.html
- https://go.dev/doc/install
- https://go.dev/doc/modules/gomod-ref
