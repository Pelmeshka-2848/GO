# Multiple Imports

## Суть

Если нужно подключить несколько пакетов, в Go используют блочный импорт.

```go
import (
	"bufio"
	"fmt"
	"os"
)
```

Это предпочтительнее, чем несколько одиночных импортов подряд.

## Пример

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
	fmt.Println(scanner.Text())
}
```

## Связанные заметки

- [[01_Theory/99_Atomic_Go/005_Minimal_Go_Program]]

