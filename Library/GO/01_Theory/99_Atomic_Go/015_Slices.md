# Slices

## Суть

`slice` — динамическая последовательность элементов одного типа.

```go
banks := []Bank{
	{Name: "Lunar Bank", Prefix: "4000"},
	{Name: "Mars Credit Union", Prefix: "5000"},
}
```

Тип `[]Bank` означает slice значений типа `Bank`.

## Отличие от массива

Массив имеет фиксированную длину:

```go
var arr [3]int
```

Slice обычно удобнее для коллекций переменной длины:

```go
items := []int{1, 2, 3}
```

## Связанные заметки

- [[01_Theory/99_Atomic_Go/014_Structs]]
- [[01_Theory/99_Atomic_Go/023_range_Loop]]

