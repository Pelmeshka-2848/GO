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

# Linux Environment Problems

## Симптомы

```text
permission denied
command not found
address already in use
no such file or directory
```

## Где возникло

- проект: Go-лабораторные в Linux, WSL или удаленной среде;
- команда: запуск приложения, работа с файлами, запуск сервера;
- окружение: shell, PATH, права файлов, system packages.

## Причина

Типовые причины: команда не установлена или не находится в PATH, файл не имеет прав на выполнение, порт занят другим процессом, путь указан относительно другой рабочей директории.

## Решение

```bash
go version
pwd
ls
```

Что проверить:

- текущую директорию;
- наличие нужного файла;
- права на файл;
- занятые порты;
- переменные окружения для Go.

## Как не повторить

- Всегда фиксировать рабочую директорию перед запуском команды.
- Использовать относительные пути от корня проекта.
- Проверять установку Go и нужных CLI-инструментов до лабораторной.

## Связанная теория

- [[01_Theory/99_Atomic_Go/000_Go_Toolchain]]
- [[01_Theory/02_Go_Tooling/Go_Run]]
- [[02_Practice/03_Debug_Diary/Port_Already_In_Use]]
