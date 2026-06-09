# 03 — Код

## Итоговое решение

```go
package main

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

func main() {
}
```

## Примечание

Код соответствует критериям проверки курса. Где требовалось, решение дополнено более чистым стилем для дальнейшего использования в базе знаний.

## Связанная теория

- [[01_Theory/99_Atomic_Go/024_Luhn_Algorithm]]
- [[01_Theory/99_Atomic_Go/025_String_Indexing_Bytes]]
- [[01_Theory/99_Atomic_Go/026_ASCII_Digit_Conversion]]
- [[01_Theory/99_Atomic_Go/027_Reverse_For_Loop]]
- [[01_Theory/99_Atomic_Go/028_Modulo_Operator]]

