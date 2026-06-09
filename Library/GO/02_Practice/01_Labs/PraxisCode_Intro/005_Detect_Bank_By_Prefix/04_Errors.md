# 04 — Ошибки

## Фактические ошибки

Ошибок не было. Важный предотвращённый дефект: не был использован возврат `&bank` из цикла `for _, bank := range banks`, потому что `bank` является копией элемента.

## Связанные записи Debug Diary

- [[02_Practice/03_Debug_Diary/Go/001_Range_Copy_Address]]

## Вывод

Ошибки и потенциальные gotcha зафиксированы отдельно только там, где они имеют повторную практическую ценность.

## Связанная теория

- [[01_Theory/99_Atomic_Go/014_Structs]]
- [[01_Theory/99_Atomic_Go/015_Slices]]
- [[01_Theory/99_Atomic_Go/019_Functions]]
- [[01_Theory/99_Atomic_Go/020_Pointers]]
- [[01_Theory/99_Atomic_Go/021_nil]]
- [[01_Theory/99_Atomic_Go/022_strings_HasPrefix]]
- [[01_Theory/99_Atomic_Go/023_range_Loop]]

