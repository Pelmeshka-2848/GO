# Pointers

## Суть

Указатель хранит адрес значения.

```go
var bank *Bank
```

Тип `*Bank` означает указатель на `Bank`.

## Получить адрес

```go
bank := Bank{Name: "Lunar Bank", Prefix: "4000"}
ptr := &bank
```

Оператор `&` берёт адрес значения.

## Зачем указатель в поиске

Функция может вернуть `*Bank`, если банк найден, или `nil`, если результата нет.

```go
func DetectBank(cardNumber string, banks []Bank) *Bank {
	return nil
}
```

## Связанные заметки

- [[01_Theory/99_Atomic_Go/021_nil]]
- [[01_Theory/99_Atomic_Go/023_range_Loop]]

