# Прочитать одну строку из stdin

## Задача

Получить строку, которую пользователь ввёл в терминале или передал через pipe.

## Шаблон

```go
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()
text := scanner.Text()
```

## Продакшен-вариант

```go
scanner := bufio.NewScanner(os.Stdin)

if scanner.Scan() {
	text := scanner.Text()
	fmt.Println(text)
}

if err := scanner.Err(); err != nil {
	fmt.Println("Ошибка чтения:", err)
}
```

## Связанная теория

- [[01_Theory/99_Atomic_Go/008_Standard_Input]]
- [[01_Theory/99_Atomic_Go/009_bufio_Scanner]]

