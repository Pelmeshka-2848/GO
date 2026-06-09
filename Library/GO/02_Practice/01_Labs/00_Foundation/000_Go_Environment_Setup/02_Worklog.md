# 02 — Ход выполнения

## 1. Установка Go SDK

Для установки был использован официальный дистрибутив Go для Windows.

Рекомендуемый вариант для Windows:

```text
go1.xx.x.windows-amd64.msi
```

После установки терминал PowerShell был перезапущен, чтобы обновилась переменная `PATH`.

## 2. Проверка версии Go

Команда:

```powershell
go version
```

Ожидаемый результат:

```text
go version go1.xx.x windows/amd64
```

Фактический результат:

```text
# Вставить вывод команды после выполнения
```

## 3. Проверка пути до go.exe

Команда:

```powershell
where.exe go
```

Ожидаемый результат:

```text
C:\Program Files\Go\bin\go.exe
```

Фактический результат:

```text
# Вставить вывод команды после выполнения
```

## 4. Проверка переменных окружения Go

Команда:

```powershell
go env GOROOT GOPATH GOMODCACHE GOPROXY GOOS GOARCH
```

Назначение параметров:

| Параметр | Назначение |
|---|---|
| `GOROOT` | путь к установленному Go SDK |
| `GOPATH` | рабочая область Go для кэша, бинарников и старого GOPATH-подхода |
| `GOMODCACHE` | кэш скачанных модулей |
| `GOPROXY` | прокси для загрузки Go-модулей |
| `GOOS` | целевая операционная система |
| `GOARCH` | целевая архитектура процессора |

Фактический результат:

```text
# Вставить вывод команды после выполнения
```

## 5. Создание рабочей директории

Команды:

```powershell
mkdir C:\GoLabs
cd C:\GoLabs
mkdir 000_Go_Environment_Setup
cd 000_Go_Environment_Setup
```

Итоговая директория проекта:

```text
C:\GoLabs\000_Go_Environment_Setup
```

## 6. Инициализация Go-модуля

Команда:

```powershell
go mod init example.com/sheve/hello-go
```

Ожидаемый результат:

```text
go: creating new go.mod: module example.com/sheve/hello-go
```

После выполнения команды появился файл:

```text
go.mod
```

## 7. Создание файла main.go

Команда для быстрого создания файла через PowerShell:

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

## 8. Форматирование кода

Команда:

```powershell
go fmt ./...
```

Назначение:

- привести код к стандартному стилю Go;
- проверить, что проект корректно распознаётся Go toolchain.

Ожидаемый результат:

```text
main.go
```

Если файл уже был отформатирован, вывод может быть пустым.

## 9. Запуск программы

Команда:

```powershell
go run .
```

Ожидаемый результат:

```text
Go environment is ready
OS: windows
Arch: amd64
```

Фактический результат:

```text
# Вставить вывод команды после выполнения
```

## 10. Добавление минимального теста

Команда:

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

## 11. Запуск тестов

Команда:

```powershell
go test ./...
```

Ожидаемый результат:

```text
ok      example.com/sheve/hello-go    0.XXXs
```

Фактический результат:

```text
# Вставить вывод команды после выполнения
```

## 12. Сборка проекта

Команды:

```powershell
mkdir bin
go build -o bin\hello-go.exe .
.\bin\hello-go.exe
```

Ожидаемый результат после запуска бинарника:

```text
Go environment is ready
OS: windows
Arch: amd64
```

## 13. Настройка GoLand

В GoLand необходимо проверить:

```text
File → Settings → Go → GOROOT
```

Ожидаемое состояние:

- GoLand видит установленный SDK;
- путь SDK указывает на установленный Go;
- проект открыт как Go-модуль;
- файл `go.mod` распознан IDE;
- запуск `main.go` доступен через Run Configuration.

## 14. Итоговая структура проекта

```text
000_Go_Environment_Setup/
├── bin/
│   └── hello-go.exe
├── go.mod
├── main.go
└── main_test.go
```

## 15. Вывод по ходу выполнения

В ходе выполнения была подготовлена минимальная среда разработки Go. Была проверена доступность Go CLI, создан Go-модуль, написана минимальная программа, добавлен базовый тест и выполнена сборка исполняемого файла.
