---
type: lab
area: Go
lab_id: 000
title: Go Environment Setup
status: ready-to-run
created: 2026-06-06
tags:
  - go
  - golang
  - environment
  - setup
  - goland
  - windows
---

# 000 — Go Environment Setup

## Кратко

Эта лабораторная работа фиксирует первичную настройку окружения для разработки на Go: установку Go SDK, проверку переменных окружения, настройку GoLand, создание первого модуля и запуск тестового проекта.

Лаба нужна как стартовая точка для всех следующих практических кейсов по Go.

## Цель

Подготовить локальное окружение разработки на Go и подтвердить, что:

- команда `go` доступна из терминала;
- Go SDK корректно установлен;
- GoLand видит Go SDK;
- можно создать Go-модуль через `go mod init`;
- можно запустить, отформатировать, протестировать и собрать минимальный Go-проект.

## Ожидаемый результат

В результате выполнения должны быть получены:

- установленный Go SDK;
- рабочая команда `go version`;
- рабочая команда `go env`;
- созданный проект `hello-go`;
- файл `go.mod`;
- файл `main.go`;
- успешный запуск `go run .`;
- успешное форматирование через `go fmt ./...`;
- успешная проверка через `go test ./...`;
- успешная сборка бинарного файла через `go build`.

## Технологии

| Компонент | Назначение |
|---|---|
| Go SDK | Компилятор, стандартная библиотека и CLI-инструменты Go |
| PowerShell | Терминал для выполнения команд |
| GoLand | Основная IDE для разработки |
| Git | Контроль версий для будущих Go-проектов |
| Windows | Основная локальная ОС |

## Папка лабораторной

```text
02_Practice/01_Labs/000_Go_Environment_Setup/
├── 00_Lab Card.md
├── 01_Task.md
├── 02_Worklog.md
├── 03_Code.md
├── 04_Errors.md
└── 05_Result.md
```

## Связанная теория

- [[01_Theory/Go/000_Go_Toolchain]]
- [[01_Theory/Go/001_Go_Modules]]
- [[01_Theory/Go/002_Go_Project_Structure]]
- [[01_Theory/Go/003_Go_CLI_Commands]]
- [[01_Theory/Go/004_GoLand_Setup]]

## Связанные практические заметки

Пока нет. Эта лаба является базовой.

## Источники

- Go — Download and install: https://go.dev/doc/install
- Go — All releases: https://go.dev/dl/
- Go — Get started with Go: https://go.dev/doc/tutorial/getting-started
- Go — How to Write Go Code: https://go.dev/doc/code
- Go — go.mod file reference: https://go.dev/doc/modules/gomod-ref
