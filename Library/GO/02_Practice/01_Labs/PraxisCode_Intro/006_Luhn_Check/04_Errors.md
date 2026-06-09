# 04 — Ошибки

## Фактические ошибки

Ошибок не было. Важный предотвращённый дефект: преобразование `ch - '0'` выполняется только после проверки, что `ch` является ASCII-цифрой.

## Связанные записи Debug Diary

- [[02_Practice/03_Debug_Diary/Go/002_ASCII_Digit_Conversion_Without_Check]]

## Вывод

Ошибки и потенциальные gotcha зафиксированы отдельно только там, где они имеют повторную практическую ценность.

## Связанная теория

- [[01_Theory/99_Atomic_Go/024_Luhn_Algorithm]]
- [[01_Theory/99_Atomic_Go/025_String_Indexing_Bytes]]
- [[01_Theory/99_Atomic_Go/026_ASCII_Digit_Conversion]]
- [[01_Theory/99_Atomic_Go/027_Reverse_For_Loop]]
- [[01_Theory/99_Atomic_Go/028_Modulo_Operator]]

