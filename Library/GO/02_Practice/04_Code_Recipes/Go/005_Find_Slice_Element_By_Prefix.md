# Найти элемент slice по префиксу и вернуть указатель

## Задача

Найти первый элемент slice, который подходит под условие, и вернуть указатель на него.

## Шаблон

```go
func FindItem(items []Item) *Item {
	for i := range items {
		if condition(items[i]) {
			return &items[i]
		}
	}

	return nil
}
```

## Пример

```go
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
```

## Важный нюанс

Если нужно вернуть указатель на элемент slice, используй индекс:

```go
return &banks[i]
```

Не используй адрес переменной цикла:

```go
return &bank
```

## Связанная теория

- [[01_Theory/99_Atomic_Go/020_Pointers]]
- [[01_Theory/99_Atomic_Go/021_nil]]
- [[01_Theory/99_Atomic_Go/022_strings_HasPrefix]]
- [[01_Theory/99_Atomic_Go/023_range_Loop]]

