# Lab 007 - Clean Architecture API — Result

## Проверка

- [ ] Папки и слои имеют понятные ответственности.
- [ ] HTTP-слой не содержит SQL.
- [ ] Бизнес-логика отделена от транспорта.
- [ ] Repository скрывает детали хранения.

## Итог

В ходе выполнения работы был разобран переход от простого REST API к более чистой архитектуре с разделением handler, service, repository и domain.

## Что стало понятнее

- Clean Architecture — это в первую очередь управление зависимостями.
- Слои нужны не для красоты, а чтобы изменения не ломали весь проект.
- Слишком раннее усложнение структуры тоже становится проблемой.

## Что повторить

- [ ] [[01_Theory/08_Architecture/Clean_Architecture]]
- [ ] [[01_Theory/08_Architecture/Layered_Architecture]]
- [ ] [[01_Theory/08_Architecture/Dependency_Injection]]
