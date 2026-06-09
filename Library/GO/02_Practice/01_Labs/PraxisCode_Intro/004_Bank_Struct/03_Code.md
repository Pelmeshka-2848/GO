# 03 — Код

## Итоговое решение

```go
package main

import "fmt"

type Bank struct {
	Name   string
	Prefix string
}

var banks = []Bank{
	{Name: "Lunar Bank", Prefix: "4000"},
	{Name: "Mars Credit Union", Prefix: "5000"},
	{Name: "Venus Express", Prefix: "6000"},
	{Name: "Saturn Ring Financial", Prefix: "7000"},
	{Name: "Jupiter Trust", Prefix: "8000"},
}

func main() {
	fmt.Println(banks)
}
```

## Примечание

Код соответствует критериям проверки курса. Где требовалось, решение дополнено более чистым стилем для дальнейшего использования в базе знаний.

## Связанная теория

- [[01_Theory/99_Atomic_Go/014_Structs]]
- [[01_Theory/99_Atomic_Go/015_Slices]]
- [[01_Theory/99_Atomic_Go/016_Composite_Literals]]
- [[01_Theory/99_Atomic_Go/017_Exported_And_Unexported_Names]]
- [[01_Theory/99_Atomic_Go/018_Package_Level_Variables]]

