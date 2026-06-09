# Получить текущее время в формате часы:минуты

## Задача

Получить текущее время и вывести его в формате `HH:MM`.

## Шаблон

```go
now := time.Now()
formatted := now.Format("15:04")
```

## Пример

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	formatted := now.Format("15:04")

	fmt.Println(formatted)
}
```

## Связанная теория

- [[01_Theory/99_Atomic_Go/011_Time_Package]]
- [[01_Theory/99_Atomic_Go/012_Time_Format_Layout]]

