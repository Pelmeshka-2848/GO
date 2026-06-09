# 03 — Код

## Итоговое решение

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	name := scanner.Text()
	fmt.Println("Привет,", name)
}
```

## Примечание

Код соответствует критериям проверки курса. Где требовалось, решение дополнено более чистым стилем для дальнейшего использования в базе знаний.

## Связанная теория

- [[01_Theory/99_Atomic_Go/006_fmt_Package]]
- [[01_Theory/99_Atomic_Go/008_Standard_Input]]
- [[01_Theory/99_Atomic_Go/009_bufio_Scanner]]
- [[01_Theory/99_Atomic_Go/010_Multiple_Imports]]

