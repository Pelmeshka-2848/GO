# Нельзя возвращать адрес переменной range при поиске элемента slice

## Ситуация

Нужно найти банк в slice и вернуть указатель на него.

## Неправильный вариант

```go
for _, bank := range banks {
	if strings.HasPrefix(cardNumber, bank.Prefix) {
		return &bank
	}
}
```

## Причина ошибки

Переменная `bank` внутри `for _, bank := range banks` является копией элемента slice. Возврат `&bank` возвращает адрес копии, а не адрес настоящего элемента внутри `banks`.

## Правильный вариант

```go
for i := range banks {
	if strings.HasPrefix(cardNumber, banks[i].Prefix) {
		return &banks[i]
	}
}
```

## Вывод

Если нужно вернуть указатель на элемент slice, следует перебирать slice по индексу.

## Связанная теория

- [[01_Theory/99_Atomic_Go/015_Slices]]
- [[01_Theory/99_Atomic_Go/020_Pointers]]
- [[01_Theory/99_Atomic_Go/023_range_Loop]]

