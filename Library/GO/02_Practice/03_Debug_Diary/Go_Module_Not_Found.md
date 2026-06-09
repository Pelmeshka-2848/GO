# Ошибка импорта локального модуля

## Контекст

Возникает при запуске Go-команд, когда проект не видит нужный модуль, пакет или корень `go.mod`.

## Текст ошибки

```text
go.mod file not found
no required module provides package
package ... is not in std
```

## Причина

Команда запущена не из корня модуля, зависимость не добавлена в `go.mod`, import path указан неверно или проект открыт не той папкой.

## Решение

- Перейти в корень проекта с `go.mod`.
- Проверить module path.
- Выполнить `go mod tidy`.
- Исправить import path на путь внутри текущего модуля.

## Как не повторить

- Один учебный проект — один понятный корень с `go.mod`.
- Не запускать Go-команды из случайной вложенной папки.
- После перемещений файлов проверять imports.

## Связанная практика

- [[02_Practice/01_Labs/00_Foundation/001_Go_Modules_And_Project_Structure/00_Lab Card]]
- [[01_Theory/02_Go_Tooling/Go_Modules]]
- [[01_Theory/02_Go_Tooling/Project_Structure]]
