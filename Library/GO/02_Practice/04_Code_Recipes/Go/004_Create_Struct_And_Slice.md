# Создать struct и slice структур

## Задача

Описать сущность и создать список таких сущностей.

## Шаблон

```go
type Entity struct {
	Field1 string
	Field2 string
}

items := []Entity{
	{Field1: "value1", Field2: "value2"},
	{Field1: "value3", Field2: "value4"},
}
```

## Пример

```go
type Bank struct {
	Name   string
	Prefix string
}

banks := []Bank{
	{Name: "Lunar Bank", Prefix: "4000"},
	{Name: "Mars Credit Union", Prefix: "5000"},
}
```

## Связанная теория

- [[01_Theory/99_Atomic_Go/014_Structs]]
- [[01_Theory/99_Atomic_Go/015_Slices]]
- [[01_Theory/99_Atomic_Go/016_Composite_Literals]]

