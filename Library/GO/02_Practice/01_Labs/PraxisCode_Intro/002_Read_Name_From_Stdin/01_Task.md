# 01 — Задание

## Тема

Прочитай имя из stdin

## Цель

Научиться читать строку из стандартного ввода `stdin` с помощью `bufio.Scanner` и использовать введённое значение в выводе программы.

## Формулировка задания


В программе был создан сканер:

```go
scanner := bufio.NewScanner(os.Stdin)
```

Нужно было вызвать `scanner.Scan()`, получить введённую строку через `scanner.Text()` и вывести приветствие с этим значением.


## Критерии проверки

- В `func main()` вызывается `scanner.Scan()`.
- Используется `scanner.Text()`, и его результат попадает в вывод программы.
- Вывод содержит `Привет` и значение из `scanner.Text()`.

## Связанная теория

- [[01_Theory/99_Atomic_Go/006_fmt_Package]]
- [[01_Theory/99_Atomic_Go/008_Standard_Input]]
- [[01_Theory/99_Atomic_Go/009_bufio_Scanner]]
- [[01_Theory/99_Atomic_Go/010_Multiple_Imports]]

