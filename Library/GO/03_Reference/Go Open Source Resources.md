---
type: reference
area: reference
status: active
tags:
  - reference
  - go
  - open-source
---

# Go Open Source Resources

Справочная заметка по внешним Go-репозиториям, которые полезны для учебного vault. Это не задания на копирование структуры или кода, а источники для сверки архитектурных решений, поиска библиотек и понимания production-подходов.

## golang-standards/project-layout

Источник: https://github.com/golang-standards/project-layout

Назначение:

- справочник по распространенным директориям в Go-проектах;
- пример структуры для крупных application-проектов;
- источник терминов вроде `cmd`, `internal`, `pkg`, `api`, `configs`, `scripts`, `deployments`.

Важно:

- это не официальный стандарт Go;
- для первых лабораторных и PoC такая структура может быть избыточной;
- начинать лучше с `go.mod` и `main.go`, а папки добавлять по мере роста проекта;
- `internal` имеет особое значение: Go ограничивает импорт пакетов из `internal` за пределами родительского дерева;
- `cmd` полезен для точек входа приложений;
- `pkg` стоит использовать осторожно, только если код реально предполагается как переиспользуемая библиотека.

Связанные заметки:

- [[01_Theory/02_Go_Tooling/Project_Structure]]
- [[01_Theory/90_Rules/Project Layout Rules]]
- [[01_Theory/08_Architecture/Layered_Architecture]]
- [[02_Practice/01_Labs/00_Foundation/001_Go_Modules_And_Project_Structure/00_Lab Card]]

## JordanMarcelino/learn-go-microservices

Источник: https://github.com/JordanMarcelino/learn-go-microservices

Назначение:

- учебный пример e-commerce системы на Go microservices;
- пример разделения на сервисы: auth, product, order, mail, gateway;
- пример инфраструктурных папок: `deployments`, `infra`, `postman`, `pkg`;
- пример тем для продвинутого разбора: API gateway, event bus, PostgreSQL, Redis, Kafka, RabbitMQ, observability.

Как использовать в курсе:

- смотреть как reference после базового REST API и PostgreSQL;
- использовать для архитектурного разбора, а не для копирования кода;
- сравнивать границы сервисов с лабораторной [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/009_Microservice_Split/00_Lab Card]];
- выделять отдельные темы в теорию: gateway, event-driven architecture, observability, tracing.

Связанные заметки:

- [[01_Theory/10_Microservices/00_Index]]
- [[01_Theory/10_Microservices/API Gateway]]
- [[01_Theory/10_Microservices/Event Driven Architecture]]
- [[01_Theory/10_Microservices/Message Broker]]
- [[01_Theory/10_Microservices/Metrics and Observability]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/012_Microservices_Basics/00_Lab Card]]

## avelino/awesome-go

Источник: https://github.com/avelino/awesome-go

Назначение:

- каталог Go-библиотек, фреймворков, инструментов и проектов;
- точка входа для поиска mature-пакетов;
- справочник, когда нужно выбрать библиотеку для конкретной категории: HTTP, database, testing, logging, configuration, CLI, observability.

Как использовать:

- не добавлять библиотеку только потому, что она есть в списке;
- сначала понять стандартную библиотеку и задачу;
- проверить активность проекта, документацию, license, количество зависимостей и простоту интеграции;
- для учебных лабораторных предпочитать стандартную библиотеку, если она закрывает задачу.

Связанные заметки:

- [[03_Reference/Links]]
- [[01_Theory/04_Stdlib/Stdlib Index]]
- [[01_Theory/11_Advanced_Go/Testing]]
- [[01_Theory/09_Production/Logging]]
- [[01_Theory/09_Production/Observability]]

## Правило использования внешних репозиториев

- Сначала понять принцип на маленьком учебном примере.
- Потом смотреть внешний репозиторий как reference.
- Не копировать структуру целиком без причины.
- Любую найденную идею переносить в vault как теорию, практический кейс, debug-запись или code recipe.
