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

# Windows Environment Problems

## Симптомы

```text
'go' is not recognized
Access is denied
The process cannot access the file because it is being used by another process
```

## Где возникло

- проект: Go-лабораторные на Windows;
- команда: `go run`, `go test`, запуск сервера, работа с файлами;
- окружение: PowerShell, Go SDK, GoLand, PATH.

## Причина

Проблемы Windows-окружения чаще всего связаны с PATH, правами доступа, занятыми портами, блокировкой файлов редактором или различиями между PowerShell и Linux-командами.

## Решение

```bash
go version
go env
```

Что проверить:

- установлен ли Go SDK;
- доступна ли команда `go` из терминала;
- не занят ли порт другим процессом;
- не открыт ли файл в программе, которая блокирует запись;
- запускается ли команда из корня проекта.

## Как не повторить

- Проверять окружение перед началом новой лабораторной.
- Фиксировать рабочие команды именно для PowerShell.
- Не хранить сборочные артефакты и временные файлы в vault.

## Связанная теория

- [[01_Theory/99_Atomic_Go/000_Go_Toolchain]]
- [[01_Theory/02_Go_Tooling/Go_Installation]]
- [[02_Practice/03_Debug_Diary/Port_Already_In_Use]]
