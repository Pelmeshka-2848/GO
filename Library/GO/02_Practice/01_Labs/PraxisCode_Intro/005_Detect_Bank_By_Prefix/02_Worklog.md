# 02 — Ход выполнения

## Выполненные шаги

1. Была объявлена функция `DetectBank` с возвратом `*Bank`.
2. Была добавлена проверка пустого `cardNumber`.
3. Был подключён пакет `strings`.
4. Поиск выполнялся через цикл `for i := range banks`.
5. Для каждого банка проверялось начало номера карты через `strings.HasPrefix`.
6. При совпадении возвращался адрес элемента slice: `&banks[i]`.
7. Если совпадений не было, функция возвращала `nil`.

## Итог выполнения

Функция определяет банк по префиксу номера карты или возвращает `nil`, если совпадения нет.

## Связанная теория

- [[01_Theory/99_Atomic_Go/014_Structs]]
- [[01_Theory/99_Atomic_Go/015_Slices]]
- [[01_Theory/99_Atomic_Go/019_Functions]]
- [[01_Theory/99_Atomic_Go/020_Pointers]]
- [[01_Theory/99_Atomic_Go/021_nil]]
- [[01_Theory/99_Atomic_Go/022_strings_HasPrefix]]
- [[01_Theory/99_Atomic_Go/023_range_Loop]]

