---
type: theory
area: theory
status: active
created: 2026-06-06
updated: 2026-06-09
tags:
  - area/theory
  - architecture
---

# Service Layer

## Кратко

Service layer содержит бизнес-сценарии приложения: создать пользователя, завершить задачу, выдать токен, проверить доступ.

## Зачем нужно

Service layer отделяет правила приложения от HTTP и базы данных.

## Что делает service

- валидирует бизнес-условия;
- координирует несколько repository;
- запускает транзакционный сценарий;
- вызывает внешние клиенты;
- возвращает результат или ошибку предметной области.

## Что не должен делать service

- читать `http.Request`;
- писать HTTP-ответ;
- собирать SQL-запросы;
- знать детали router-а;
- логировать каждую мелкую ошибку без причины.

## Правила

- Методы service обычно принимают context.
- Входные данные должны быть уже разобраны transport-слоем.
- Ошибки service должны быть понятны handler-у для преобразования в HTTP.
- Service зависит от абстракций repository/client, если это помогает тестировать.

## Типичные ошибки

| Ошибка | Почему возникает | Как исправить |
|---|---|---|
| Service принимает `http.ResponseWriter` | Смешение transport и logic | Передавать обычные структуры |
| Вся логика в repository | Repository стал service | Вернуть сценарий в service |
| Нет service вообще | Маленький проект вырос | Вынести сценарии постепенно |

## Связанные темы

- [[01_Theory/08_Architecture/Layered Architecture]]
- [[01_Theory/06_Databases/Repository_Layer]]
- [[01_Theory/08_Architecture/Dependency Injection]]

