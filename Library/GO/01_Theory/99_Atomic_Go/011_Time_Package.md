# time Package

## Суть

`time` — стандартный пакет Go для работы с датой, временем, длительностями и таймерами.

## Получить текущее время

```go
now := time.Now()
```

`time.Now()` возвращает текущее локальное время.

## Пример

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
}
```

## Связанные заметки

- [[01_Theory/99_Atomic_Go/012_Time_Format_Layout]]

