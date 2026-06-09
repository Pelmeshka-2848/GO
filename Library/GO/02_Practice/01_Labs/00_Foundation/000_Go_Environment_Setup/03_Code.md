# 03 — Код

## Структура проекта

```text
000_Go_Environment_Setup/
├── bin/
│   └── hello-go.exe
├── go.mod
├── main.go
└── main_test.go
```

## go.mod

Файл `go.mod` создаётся командой:

```powershell
go mod init example.com/sheve/hello-go
```

Пример содержимого:

```go
module example.com/sheve/hello-go

go 1.xx.x
```

Примечание:

- точная версия в строке `go` зависит от установленной версии Go;
- файл `go.mod` описывает модуль и минимальную версию Go для проекта;
- в дальнейшем сюда будут добавляться зависимости.

## main.go

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("Go environment is ready")
    fmt.Println("OS:", runtime.GOOS)
    fmt.Println("Arch:", runtime.GOARCH)
}
```

## main_test.go

```go
package main

import (
    "runtime"
    "testing"
)

func TestRuntimeValuesAreAvailable(t *testing.T) {
    if runtime.GOOS == "" {
        t.Fatal("GOOS is empty")
    }

    if runtime.GOARCH == "" {
        t.Fatal("GOARCH is empty")
    }
}
```

## Команды создания проекта

```powershell
mkdir C:\GoLabs
cd C:\GoLabs
mkdir 000_Go_Environment_Setup
cd 000_Go_Environment_Setup

go mod init example.com/sheve/hello-go
```

## Команды создания файлов через PowerShell

### main.go

```powershell
@'
package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("Go environment is ready")
    fmt.Println("OS:", runtime.GOOS)
    fmt.Println("Arch:", runtime.GOARCH)
}
'@ | Set-Content -Encoding UTF8 main.go
```

### main_test.go

```powershell
@'
package main

import (
    "runtime"
    "testing"
)

func TestRuntimeValuesAreAvailable(t *testing.T) {
    if runtime.GOOS == "" {
        t.Fatal("GOOS is empty")
    }

    if runtime.GOARCH == "" {
        t.Fatal("GOARCH is empty")
    }
}
'@ | Set-Content -Encoding UTF8 main_test.go
```

## Команды проверки

```powershell
go fmt ./...
go run .
go test ./...
```

## Команды сборки

```powershell
mkdir bin
go build -o bin\hello-go.exe .
.\bin\hello-go.exe
```

## Ожидаемый вывод программы

```text
Go environment is ready
OS: windows
Arch: amd64
```

## Ключевые элементы кода

| Элемент | Назначение |
|---|---|
| `package main` | объявляет исполняемый пакет |
| `func main()` | точка входа в программу |
| `fmt.Println` | выводит текст в консоль |
| `runtime.GOOS` | показывает целевую ОС |
| `runtime.GOARCH` | показывает архитектуру |
| `testing` | стандартный пакет для тестов |
| `go test` | запускает тесты проекта |

## Минимальный workflow Go-проекта

```text
создать директорию
→ go mod init
→ написать main.go
→ go fmt
→ go run
→ go test
→ go build
```
