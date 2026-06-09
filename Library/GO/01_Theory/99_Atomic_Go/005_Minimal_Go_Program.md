# Minimal Go Program

## Суть

Минимальная исполняемая программа на Go обычно содержит:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello")
}
```

## Основные элементы

### `package main`

`package main` указывает, что пакет собирается как исполняемая программа.

### `func main()`

`main` — точка входа. С неё начинается выполнение программы.

### `import`

`import` подключает стандартные или внешние пакеты.

## Связанные заметки

- [[01_Theory/99_Atomic_Go/006_fmt_Package]]
- [[01_Theory/99_Atomic_Go/010_Multiple_Imports]]

