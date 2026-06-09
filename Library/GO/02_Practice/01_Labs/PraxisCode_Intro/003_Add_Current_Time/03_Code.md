# 03 — Код

## Итоговое решение

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	name := scanner.Text()

	now := time.Now()
	formatted := now.Format("15:04")

	fmt.Printf("Привет, %s! Сейчас %s\n", name, formatted)
}
```

## Примечание

Код соответствует критериям проверки курса. Где требовалось, решение дополнено более чистым стилем для дальнейшего использования в базе знаний.

## Связанная теория

- [[01_Theory/99_Atomic_Go/008_Standard_Input]]
- [[01_Theory/99_Atomic_Go/009_bufio_Scanner]]
- [[01_Theory/99_Atomic_Go/010_Multiple_Imports]]
- [[01_Theory/99_Atomic_Go/011_Time_Package]]
- [[01_Theory/99_Atomic_Go/012_Time_Format_Layout]]
- [[01_Theory/99_Atomic_Go/013_fmt_Printf]]

