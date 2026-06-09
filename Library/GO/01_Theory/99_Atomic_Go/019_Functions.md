# Functions

## Суть

Функция — именованный блок кода с параметрами и возвращаемым значением.

```go
func DetectBank(cardNumber string, banks []Bank) *Bank {
	return nil
}
```

## Части сигнатуры

```go
func DetectBank(cardNumber string, banks []Bank) *Bank
```

- `DetectBank` — имя функции;
- `cardNumber string` — параметр;
- `banks []Bank` — параметр;
- `*Bank` — тип возвращаемого значения.

## Связанные заметки

- [[01_Theory/99_Atomic_Go/020_Pointers]]
- [[01_Theory/99_Atomic_Go/021_nil]]

