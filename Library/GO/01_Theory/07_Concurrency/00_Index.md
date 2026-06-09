# Concurrency Index

Раздел про конкурентность в Go: goroutines, channels, синхронизацию, отмену операций, timeouts и типичные ошибки параллельного кода.

## Рекомендуемый порядок

1. [[01_Theory/07_Concurrency/Goroutines|Goroutines]]
2. [[01_Theory/07_Concurrency/Channels|Channels]]
3. [[01_Theory/07_Concurrency/Unbuffered_Channels|Unbuffered Channels]]
4. [[01_Theory/07_Concurrency/Buffered_Channels|Buffered Channels]]
5. [[01_Theory/07_Concurrency/WaitGroup|WaitGroup]]
6. [[01_Theory/07_Concurrency/Select|Select]]
7. [[01_Theory/07_Concurrency/Timeouts|Timeouts]]
8. [[01_Theory/07_Concurrency/Context|Context]]
9. [[01_Theory/07_Concurrency/Mutex|Mutex]]
10. [[01_Theory/07_Concurrency/Data_Race|Data Race]]
11. [[01_Theory/07_Concurrency/Deadlock|Deadlock]]
12. [[01_Theory/07_Concurrency/Worker_Pool|Worker Pool]]

## Что нужно уметь

- запускать работу в goroutine;
- дожидаться завершения нескольких задач;
- передавать данные через channels;
- ограничивать параллелизм;
- отменять долгие операции через context;
- защищать shared state через mutex;
- находить data race и deadlock.

## Практика

- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/006_Concurrent_URL_Checker/00_Lab Card]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/007_Worker_Pool/00_Lab Card]]

