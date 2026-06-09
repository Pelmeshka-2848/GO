---
type: theory
area: Go
title: Go Project Structure
tags:
  - go
  - golang
  - project-structure
  - packages
  - cmd
  - internal
  - pkg
---

# Go Project Structure

## Кратко

В Go структура проекта строится вокруг:

- модуля;
- пакетов;
- файла `go.mod`;
- директорий с `.go`-файлами;
- правил видимости пакетов.

Минимальный проект:

```text
hello-go/
├── go.mod
└── main.go
```

## Минимальный исполняемый проект

```text
hello-go/
├── go.mod
├── main.go
└── main_test.go
```

Файл `main.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go")
}
```

Ключевой признак исполняемой программы:

```go
package main
```

и наличие функции:

```go
func main()
```

## Пакет

Пакет — базовая единица организации кода в Go.

В начале каждого `.go`-файла указывается пакет:

```go
package main
```

или:

```go
package user
```

Правило:

```text
в одной директории все обычные .go-файлы должны принадлежать одному пакету
```

Пример:

```text
user/
├── user.go       package user
├── service.go    package user
└── user_test.go  package user
```

## package main

`package main` используется для программ, которые собираются в исполняемый файл.

Пример:

```text
cmd/app/main.go
```

```go
package main

func main() {
}
```

## Обычные пакеты

Обычный пакет используется как библиотечный код.

Пример:

```text
internal/user/service.go
```

```go
package user

func CreateUser() {
}
```

Такой пакет не запускается сам по себе через `func main`, но может импортироваться другими частями проекта.

## Import path

Import path строится относительно module path.

Если `go.mod` содержит:

```go
module example.com/sheve/app
```

и есть пакет:

```text
internal/user
```

то импорт будет:

```go
import "example.com/sheve/app/internal/user"
```

## Рекомендуемая простая структура для начала

Для первых лабораторных достаточно:

```text
project/
├── go.mod
├── main.go
└── main_test.go
```

Для проекта с несколькими файлами:

```text
project/
├── go.mod
├── main.go
├── config.go
├── handlers.go
└── handlers_test.go
```

## Структура для приложения

Более взрослая структура:

```text
project/
├── go.mod
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── config/
│   ├── handler/
│   ├── service/
│   └── repository/
├── pkg/
├── migrations/
├── configs/
└── README.md
```

## cmd

`cmd` обычно содержит точки входа приложений.

Пример:

```text
cmd/
├── api/
│   └── main.go
└── worker/
    └── main.go
```

Каждая директория внутри `cmd` — отдельная программа.

Пример запуска:

```bash
go run ./cmd/api
```

Сборка:

```bash
go build -o bin/api.exe ./cmd/api
```

## internal

`internal` — специальная директория Go.

Код внутри `internal` нельзя импортировать извне родительского дерева проекта.

Пример:

```text
project/
├── go.mod
└── internal/
    └── user/
        └── service.go
```

Такой пакет можно импортировать внутри проекта, но нельзя использовать как публичную библиотеку из другого внешнего проекта.

Это полезно для:

- бизнес-логики приложения;
- внутренних сервисов;
- репозиториев;
- обработчиков;
- конфигурации.

## pkg

`pkg` часто используют для кода, который потенциально может быть импортирован снаружи.

Пример:

```text
pkg/
└── validator/
```

Важно:

- `pkg` не является специальной директорией на уровне языка;
- это соглашение, а не механизм компилятора;
- для учебных проектов можно не использовать `pkg`, пока нет явной необходимости.

## test files

Тесты в Go лежат в файлах:

```text
*_test.go
```

Пример:

```text
main_test.go
user_service_test.go
```

Тестовая функция:

```go
func TestName(t *testing.T) {
}
```

Запуск всех тестов:

```bash
go test ./...
```

## README.md

`README.md` полезен даже в учебных проектах.

Минимальное содержание:

```text
# Project Name

## Run

go run .

## Test

go test ./...

## Build

go build -o bin/app .
```

## bin

Директория `bin` часто используется для собранных бинарников:

```text
bin/
└── app.exe
```

Сборка:

```bash
go build -o bin/app.exe .
```

Обычно `bin` добавляют в `.gitignore`.

## configs

`configs` может содержать примеры конфигураций:

```text
configs/
└── config.example.yaml
```

Секреты и реальные пароли в Git не кладут.

## migrations

Для проектов с базой данных можно использовать:

```text
migrations/
├── 000001_create_users.up.sql
└── 000001_create_users.down.sql
```

В стартовой Go-лабе это не требуется.

## Что не надо делать в начале

Не стоит сразу усложнять проект:

```text
project/
├── adapters/
├── application/
├── domain/
├── infrastructure/
├── interfaces/
└── usecases/
```

Такая структура может быть полезна в DDD/Clean Architecture, но для первых Go-лабораторных она создаёт лишний шум.

Начинать лучше с простого:

```text
project/
├── go.mod
├── main.go
└── main_test.go
```

## Практическое правило

Структура должна расти от задачи, а не заранее.

```text
1 файл → несколько файлов → несколько пакетов → cmd/internal → отдельные сервисы
```

## Частые ошибки структуры

### Открыт не корень проекта

Плохо:

```text
открыт только файл main.go
```

Хорошо:

```text
открыта папка, где лежит go.mod
```

### Разные package в одной директории

Плохо:

```text
project/
├── main.go       package main
└── service.go    package service
```

Хорошо:

```text
project/
├── main.go
└── internal/
    └── service/
        └── service.go
```

### Слишком ранний pkg

Не надо выносить код в `pkg`, если он не является публичной библиотекой.

Лучше:

```text
internal/
```

## Минимальные шаблоны

### Однофайловая лабораторная

```text
lab/
├── go.mod
└── main.go
```

### Лабораторная с тестом

```text
lab/
├── go.mod
├── main.go
└── main_test.go
```

### Небольшое CLI-приложение

```text
app/
├── go.mod
├── cmd/
│   └── app/
│       └── main.go
└── internal/
    └── app/
        └── app.go
```

## Что запомнить

- Корень Go-проекта обычно содержит `go.mod`.
- `package main` нужен для исполняемого приложения.
- `func main()` — точка входа.
- Один каталог — один пакет.
- `internal` ограничивает импорт извне.
- `cmd` удобно использовать для нескольких приложений.
- `pkg` — соглашение, а не специальный механизм языка.
- Для первых лабораторных достаточно простой структуры.

## Связанные заметки

- [[01_Theory/Go/001_Go_Modules]]
- [[01_Theory/Go/003_Go_CLI_Commands]]
- [[01_Theory/Go/004_GoLand_Setup]]

## Источники

- https://go.dev/doc/modules/layout
- https://go.dev/doc/code
- https://go.dev/doc/modules/managing-source
- https://go.dev/doc/tutorial/create-module
