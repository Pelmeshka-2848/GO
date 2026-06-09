# 03 — Код

## Итоговое решение

```go
package main

import (
	"fmt"
	"strings"
)

type Bank struct {
	Name   string
	Prefix string
}

func DetectBank(cardNumber string, banks []Bank) *Bank {
	if cardNumber == "" {
		return nil
	}

	for i := range banks {
		if strings.HasPrefix(cardNumber, banks[i].Prefix) {
			return &banks[i]
		}
	}

	return nil
}

func main() {
	banks := []Bank{
		{Name: "Lunar Bank", Prefix: "4000"},
		{Name: "Mars Credit Union", Prefix: "5000"},
	}

	bank := DetectBank("4000123456789017", banks)

	if bank != nil {
		fmt.Println(bank.Name)
	} else {
		fmt.Println("Банк не найден")
	}
}
```

## Примечание

Код соответствует критериям проверки курса. Где требовалось, решение дополнено более чистым стилем для дальнейшего использования в базе знаний.

## Связанная теория

- [[01_Theory/99_Atomic_Go/014_Structs]]
- [[01_Theory/99_Atomic_Go/015_Slices]]
- [[01_Theory/99_Atomic_Go/019_Functions]]
- [[01_Theory/99_Atomic_Go/020_Pointers]]
- [[01_Theory/99_Atomic_Go/021_nil]]
- [[01_Theory/99_Atomic_Go/022_strings_HasPrefix]]
- [[01_Theory/99_Atomic_Go/023_range_Loop]]

