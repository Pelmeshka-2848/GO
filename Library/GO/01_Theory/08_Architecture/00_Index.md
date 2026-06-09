# Architecture Index

Раздел про архитектуру Go backend-приложений: слои, зависимости, DTO, сервисы, кэш, очереди, API gateway и границы ответственности.

## Рекомендуемый порядок

1. [[01_Theory/08_Architecture/Monolith|Monolith]]
2. [[01_Theory/08_Architecture/Layered Architecture|Layered Architecture]]
3. [[01_Theory/08_Architecture/Service Layer|Service Layer]]
4. [[01_Theory/08_Architecture/Domain Layer|Domain Layer]]
5. [[01_Theory/08_Architecture/DTO|DTO]]
6. [[01_Theory/08_Architecture/Dependency Injection|Dependency Injection]]
7. [[01_Theory/08_Architecture/Clean Architecture|Clean Architecture]]
8. [[01_Theory/08_Architecture/Hexagonal Architecture|Hexagonal Architecture]]
9. [[01_Theory/08_Architecture/Cache|Cache]]
10. [[01_Theory/08_Architecture/Rate_Limiting|Rate Limiting]]
11. [[01_Theory/08_Architecture/Message_Queues|Message Queues]]
12. [[01_Theory/08_Architecture/API_Gateway|API Gateway]]
13. [[01_Theory/08_Architecture/Microservices|Microservices]]

## Что нужно уметь

- отделять HTTP, business logic и persistence;
- понимать направление зависимостей;
- проектировать service layer;
- использовать DTO на границе API;
- применять dependency injection без лишней магии;
- не усложнять маленький проект раньше времени.

## Практика

- [[02_Practice/01_Labs/01_Basic_Backend_Track/007_Clean_Architecture_API/00_Lab Card]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/005_Clean_REST_API_Project/00_Lab Card]]
- [[02_Practice/02_Development_Cases/Case - REST API Skeleton]]

## Внешние ориентиры

- [[03_Reference/Go Open Source Resources|Go Open Source Resources]]
- `golang-standards/project-layout` — смотреть как справочник директорий, а не как обязательный стандарт.
- `awesome-go` — использовать для поиска библиотек после понимания задачи и стандартной библиотеки.
