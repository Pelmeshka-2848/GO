# Lab 007 - Clean Architecture API — Errors

| Ошибка | Причина | Решение | Связанная заметка |
|---|---|---|---|
| Handler содержит SQL | HTTP-слой смешан со слоем хранения | Вынести работу с базой в repository | [[01_Theory/08_Architecture/Layered_Architecture]] |
| Service зависит от `http.Request` | Бизнес-логика привязана к транспорту | Передавать в service доменные DTO или параметры | [[01_Theory/08_Architecture/DTO]] |
| Интерфейсы созданы без необходимости | Архитектура усложнена раньше требований | Оставлять интерфейс там, где есть реальная замена реализации | [[01_Theory/90_Rules/Interface Design Rules]] |
