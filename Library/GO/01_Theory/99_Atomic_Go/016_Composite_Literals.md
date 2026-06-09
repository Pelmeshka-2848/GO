# Composite Literals

## Суть

Составной литерал позволяет создать значение структуры, массива, slice или map.

## Структура

```go
Bank{Name: "Lunar Bank", Prefix: "4000"}
```

## Slice структур

```go
banks := []Bank{
	{Name: "Lunar Bank", Prefix: "4000"},
	{Name: "Mars Credit Union", Prefix: "5000"},
}
```

Внутри `[]Bank{}` тип `Bank` можно не повторять у каждого элемента.

## Связанные заметки

- [[01_Theory/99_Atomic_Go/014_Structs]]
- [[01_Theory/99_Atomic_Go/015_Slices]]

