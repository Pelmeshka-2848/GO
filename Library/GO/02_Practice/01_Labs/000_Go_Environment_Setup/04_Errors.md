# 04 — Ошибки

## Ошибка 1. Команда go не найдена

### Симптом

```powershell
go : The term 'go' is not recognized as the name of a cmdlet, function, script file, or operable program.
```

### Причина

PowerShell не видит `go.exe`. Обычно это происходит по одной из причин:

- Go SDK не установлен;
- путь `C:\Program Files\Go\bin` не добавлен в `PATH`;
- терминал был открыт до установки Go и не подхватил обновлённый `PATH`.

### Решение

1. Закрыть PowerShell.
2. Открыть PowerShell заново.
3. Проверить команду:

```powershell
go version
```

4. Если ошибка осталась, проверить наличие файла:

```powershell
dir "C:\Program Files\Go\bin\go.exe"
```

5. Если файл существует, добавить путь в `PATH` через настройки Windows.

---

## Ошибка 2. GoLand не видит Go SDK

### Симптом

В GoLand проект открыт, но IDE не подсвечивает Go-код корректно или предлагает настроить SDK.

### Причина

В GoLand не выбран путь к установленному Go SDK.

### Решение

Открыть настройки:

```text
File → Settings → Go → GOROOT
```

Указать путь:

```text
C:\Program Files\Go
```

После этого проверить, что IDE распознала версию Go.

---

## Ошибка 3. go.mod file not found

### Симптом

```text
go: go.mod file not found in current directory or any parent directory
```

### Причина

Команда `go run .`, `go test ./...` или `go build` была выполнена вне директории Go-модуля.

### Решение

Перейти в директорию проекта:

```powershell
cd C:\GoLabs\000_Go_Environment_Setup
```

Если файла `go.mod` нет, создать модуль:

```powershell
go mod init example.com/sheve/hello-go
```

---

## Ошибка 4. Несколько версий Go в системе

### Симптом

Команда `go version` показывает не ту версию Go, которая была установлена последней.

### Причина

В `PATH` может быть несколько путей к разным установкам Go.

### Диагностика

```powershell
where.exe go
```

Если вывод содержит несколько строк, Windows берёт первый найденный путь.

### Решение

- удалить старую версию Go;
- почистить `PATH`;
- оставить актуальный путь выше остальных:

```text
C:\Program Files\Go\bin
```

---

## Ошибка 5. Не запускается собранный .exe

### Симптом

```powershell
.\bin\hello-go.exe : The term '.\bin\hello-go.exe' is not recognized...
```

### Причина

Файл не был собран или команда выполняется не из корня проекта.

### Решение

Проверить наличие файла:

```powershell
dir .\bin
```

Повторить сборку:

```powershell
mkdir bin
go build -o bin\hello-go.exe .
```

Запустить:

```powershell
.\bin\hello-go.exe
```

---

## Ошибка 6. go test выводит `[no test files]`

### Симптом

```text
?       example.com/sheve/hello-go    [no test files]
```

### Причина

В проекте нет файлов с суффиксом `_test.go`.

### Решение

Создать файл:

```text
main_test.go
```

И добавить минимальный тест:

```go
package main

import "testing"

func TestSmoke(t *testing.T) {
}
```

После этого выполнить:

```powershell
go test ./...
```

---

## Ошибка 7. Проблемы с загрузкой внешних модулей

### Симптом

При добавлении внешних зависимостей команды Go долго выполняются или завершаются ошибкой сети.

### Причина

Go использует `GOPROXY` для загрузки модулей. В некоторых сетях доступ к прокси или репозиториям может быть ограничен.

### Диагностика

```powershell
go env GOPROXY
```

### Базовое решение

Оставить стандартное значение:

```powershell
go env -w GOPROXY=https://proxy.golang.org,direct
```

Если в конкретной сети это не работает, нужно отдельно фиксировать сетевые ограничения и выбранный proxy-подход.

---

## Быстрый чек-лист диагностики

```powershell
go version
where.exe go
go env GOROOT GOPATH GOMODCACHE GOPROXY GOOS GOARCH
pwd
dir
go env GOMOD
go test ./...
```
