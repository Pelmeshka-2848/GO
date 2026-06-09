# Pprof

## Что это

`pprof` — инструмент профилирования Go-приложений: CPU, memory, goroutines, blocking, mutex.

## Зачем нужно

Pprof помогает находить реальные bottleneck-и, утечки памяти, зависшие goroutines и дорогие участки кода.

## Что можно смотреть

- CPU profile;
- heap profile;
- goroutine profile;
- block profile;
- mutex profile;
- trace.

## Правила

- Профилировать нужно реальную проблему, а не “на всякий случай”.
- Сравнивай profiles до и после изменения.
- Не включай debug endpoint публично без защиты.
- Для web-сервера pprof endpoint должен быть закрыт от внешнего доступа.

## Типичные ошибки

- оптимизировать без измерений;
- смотреть profile без нагрузки;
- открыть pprof наружу;
- делать вывод по одному короткому запуску.

## Связанные темы

- [[01_Theory/11_Advanced_Go/Benchmarks]]
- [[01_Theory/09_Production/Observability]]


