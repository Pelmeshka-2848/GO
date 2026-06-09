# Вывод значения переменной через fmt.Println

## Задача

Вывести текст вместе со значением переменной.

## Шаблон

```go
fmt.Println("Текст", variable)
```

## Пример

```go
package main

import "fmt"

func main() {
	name := "Гоша"
	fmt.Println("Привет,", name)
}
```

## Связанная теория

- [[01_Theory/99_Atomic_Go/006_fmt_Package]]

