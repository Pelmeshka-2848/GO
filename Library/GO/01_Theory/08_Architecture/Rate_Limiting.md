# Rate Limiting

## Что это

Rate limiting ограничивает количество запросов или операций за период времени.

## Зачем нужно

Ограничения защищают сервис от перегрузки, brute force, случайных циклов и неравномерной нагрузки.

## Где применять

- login endpoint;
- public API;
- внешние API-клиенты;
- background jobs;
- per-user или per-IP ограничения.

## Подходы

- fixed window;
- sliding window;
- token bucket;
- leaky bucket;
- concurrency limit.

## Правила

- Ограничение должно соответствовать риску endpoint-а.
- Клиенту полезно возвращать `429 Too Many Requests`.
- Для распределенной системы локального in-memory limiter может быть недостаточно.
- Rate limit не заменяет authentication.

## Типичные ошибки

- одинаковый limit для всех endpoint-ов;
- нет ответа `429`;
- limiter только в памяти одного instance при нескольких репликах;
- не учитывать доверенные внутренние вызовы.

## Связанные темы

- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/Rate_Limits]]
- [[01_Theory/05_HTTP_Backend/02_API_Practices/Middleware]]

