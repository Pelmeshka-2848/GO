# range Loop

## Суть

`range` используется для перебора коллекций.

## Перебор по индексу

```go
for i := range banks {
	fmt.Println(banks[i])
}
```

## Перебор с копией элемента

```go
for _, bank := range banks {
	fmt.Println(bank)
}
```

## Gotcha при указателях

Если нужно вернуть указатель на элемент slice, нужен индекс:

```go
for i := range banks {
	return &banks[i]
}
```

Не нужно возвращать адрес переменной цикла:

```go
for _, bank := range banks {
	return &bank
}
```

`bank` — копия элемента.

## Связанные заметки

- [[01_Theory/99_Atomic_Go/015_Slices]]
- [[01_Theory/99_Atomic_Go/020_Pointers]]

