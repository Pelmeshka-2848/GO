# Structs

## Суть

`struct` — пользовательский тип, который группирует несколько связанных полей.

```go
type Bank struct {
	Name   string
	Prefix string
}
```

## Создание значения

```go
bank := Bank{Name: "Lunar Bank", Prefix: "4000"}
```

## Когда использовать

`struct` используют, когда данные логически относятся к одной сущности.

Например, банк имеет название и префикс карты. Эти значения удобнее держать в одном объекте, а не в отдельных переменных.

## Связанные заметки

- [[01_Theory/99_Atomic_Go/016_Composite_Literals]]
- [[01_Theory/99_Atomic_Go/017_Exported_And_Unexported_Names]]

