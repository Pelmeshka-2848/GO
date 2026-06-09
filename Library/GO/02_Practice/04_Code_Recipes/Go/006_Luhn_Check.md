# Проверить строку по алгоритму Луна

## Задача

Проверить, проходит ли номер карты контрольную сумму по алгоритму Луна.

## Шаблон

```go
func LuhnCheck(cardNumber string) bool {
	if cardNumber == "" {
		return false
	}

	sum := 0
	doubleNext := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		ch := cardNumber[i]

		if ch < '0' || ch > '9' {
			return false
		}

		digit := int(ch - '0')

		if doubleNext {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		doubleNext = !doubleNext
	}

	return sum%10 == 0
}
```

## Примеры

```go
LuhnCheck("4000123456789017") // true
LuhnCheck("79927398713")      // true
LuhnCheck("1234567890123456") // false
LuhnCheck("")                 // false
LuhnCheck("4000abc")          // false
```

## Важный нюанс

Алгоритм Луна проверяет только контрольную сумму. Для полноценного валидатора нужны дополнительные проверки: длина, допустимые символы, префикс и банк-эмитент.

## Связанная теория

- [[01_Theory/99_Atomic_Go/024_Luhn_Algorithm]]
- [[01_Theory/99_Atomic_Go/025_String_Indexing_Bytes]]
- [[01_Theory/99_Atomic_Go/026_ASCII_Digit_Conversion]]
- [[01_Theory/99_Atomic_Go/027_Reverse_For_Loop]]
- [[01_Theory/99_Atomic_Go/028_Modulo_Operator]]

