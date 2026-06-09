# Graceful Shutdown

## Что это

Graceful shutdown — корректное завершение приложения без резкого обрыва текущих запросов и фоновых задач.

## Зачем нужно

При деплое, перезапуске или остановке сервиса нужно дать текущей работе завершиться и закрыть ресурсы.

## Что закрывать

- HTTP server;
- database pool;
- workers;
- message consumers;
- log exporters;
- external clients, если они требуют закрытия.

## Правила

- Слушай системные сигналы.
- Используй context с timeout.
- Сначала перестань принимать новые запросы.
- Затем дай текущим операциям завершиться.
- После timeout заверши процесс принудительно или верни ошибку shutdown.

## Частые ошибки

- завершать процесс сразу через `os.Exit`;
- не закрывать БД-пул;
- не останавливать workers;
- делать shutdown без timeout;
- не логировать результат остановки.

## Связанные темы

- [[01_Theory/05_HTTP_Backend/02_API_Practices/Graceful Shutdown]]
- [[01_Theory/07_Concurrency/Context]]
- [[01_Theory/04_Stdlib/net-http package]]

