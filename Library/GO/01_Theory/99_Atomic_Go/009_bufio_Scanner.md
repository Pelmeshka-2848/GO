# bufio.Scanner

## Суть

`bufio.Scanner` — инструмент для чтения данных из потока по токенам. По умолчанию токеном является строка.

## Базовый шаблон

```go
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()
text := scanner.Text()
```

## Что делает `Scan`

`scanner.Scan()` продвигает сканер к следующему токену и возвращает `true`, если чтение успешно.

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

## Gotcha

В учебных задачах результат `Scan()` часто не проверяют. В реальном коде нужно проверять и `Scan()`, и `scanner.Err()`.

## Связанные заметки

- [[01_Theory/99_Atomic_Go/008_Standard_Input]]

