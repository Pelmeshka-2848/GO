# 01 — Задание

## Тема

Определение банка по префиксу

## Цель

Научиться искать элемент в slice по условию и возвращать указатель на найденный элемент.

## Формулировка задания


Нужно было реализовать функцию:

```go
func DetectBank(cardNumber string, banks []Bank) *Bank
```

Функция должна возвращать `nil` для пустого номера, искать банк по совпадению префикса через `strings.HasPrefix`, возвращать указатель на элемент slice при совпадении и `nil`, если банк не найден.


## Критерии проверки

- Функция `DetectBank(cardNumber string, banks []Bank) *Bank` объявлена.
- Для пустого `cardNumber` функция возвращает `nil`.
- Поиск использует `strings.HasPrefix`.
- При совпадении возвращается указатель на элемент slice по индексу.
- При отсутствии совпадений функция возвращает `nil`.

## Связанная теория

- [[01_Theory/99_Atomic_Go/014_Structs]]
- [[01_Theory/99_Atomic_Go/015_Slices]]
- [[01_Theory/99_Atomic_Go/019_Functions]]
- [[01_Theory/99_Atomic_Go/020_Pointers]]
- [[01_Theory/99_Atomic_Go/021_nil]]
- [[01_Theory/99_Atomic_Go/022_strings_HasPrefix]]
- [[01_Theory/99_Atomic_Go/023_range_Loop]]

