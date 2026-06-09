# Индекс лабораторных Go

## Назначение

Главный индекс лабораторных и практических треков по Go.

Канонический вход в лабораторные: [[02_Practice/01_Labs/00_Labs_Index]].

## Принцип организации

- `00_Foundation` — базовая подготовка окружения и структуры проекта.
- `01_Basic_Backend_Track` — базовая backend-линейка.
- `02_TaskFlow_Backend_Track` — расширенный backend-трек TaskFlow.
- `PraxisCode_Intro` — отдельный курс PraxisCode Intro с собственной нумерацией.

## 00 Foundation

Базовый трек подготовки окружения и структуры Go-проекта.

- [[02_Practice/01_Labs/00_Foundation/000_Go_Environment_Setup/00_Lab Card|000 — Настройка окружения Go]]
- [[02_Practice/01_Labs/00_Foundation/001_Go_Modules_And_Project_Structure/00_Lab Card|001 — Go modules и структура проекта]]

## 01 Basic Backend Track

Базовая backend-линейка: CLI, HTTP, REST, PostgreSQL, JWT и Clean Architecture.

- [[02_Practice/01_Labs/01_Basic_Backend_Track/002_CLI_Calculator/00_Lab Card|002 — CLI-калькулятор]]
- [[02_Practice/01_Labs/01_Basic_Backend_Track/003_HTTP_Server/00_Lab Card|003 — HTTP-сервер]]
- [[02_Practice/01_Labs/01_Basic_Backend_Track/004_REST_CRUD_API/00_Lab Card|004 — REST CRUD API]]
- [[02_Practice/01_Labs/01_Basic_Backend_Track/005_PostgreSQL_CRUD/00_Lab Card|005 — PostgreSQL CRUD]]
- [[02_Practice/01_Labs/01_Basic_Backend_Track/006_Auth_JWT/00_Lab Card|006 — Авторизация через JWT]]
- [[02_Practice/01_Labs/01_Basic_Backend_Track/007_Clean_Architecture_API/00_Lab Card|007 — Clean Architecture API]]

## 02 TaskFlow Backend Track

Трек разработки TaskFlow backend: REST API, PostgreSQL, concurrency, JWT, тестирование, production-ready подход и микросервисы.

- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/002_Markdown_HTML_Converter/00_Lab Card|002 — Markdown to HTML Converter]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/003_REST_API_Todo_Service/00_Lab Card|003 — REST API Todo Service]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/004_REST_API_With_PostgreSQL/00_Lab Card|004 — REST API с PostgreSQL]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/005_Clean_REST_API_Project/00_Lab Card|005 — Чистая структура REST API]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/006_Concurrent_URL_Checker/00_Lab Card|006 — Concurrent URL Checker]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/007_Worker_Pool/00_Lab Card|007 — Worker Pool]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/008_Auth_API_With_JWT/00_Lab Card|008 — Auth API with JWT]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/009_Microservice_Split/00_Lab Card|009 — Разделение на микросервисы]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/010_Testing_REST_API/00_Lab Card|010 — Тестирование REST API]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/011_Production_Ready_API/00_Lab Card|011 — Production-ready API]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/012_Microservices_Basics/00_Lab Card|012 — Основы микросервисов]]
- [[02_Practice/01_Labs/02_TaskFlow_Backend_Track/013_Final_TaskFlow_Backend/00_Lab Card|013 — Финальный TaskFlow backend]]

## PraxisCode Intro

Практический курс PraxisCode Intro вынесен в отдельную ветку, чтобы не смешивать его нумерацию с другими треками.

- [[02_Practice/01_Labs/PraxisCode_Intro/000_Index|PraxisCode Intro — индекс]]


## Правило нумерации

Нумерация лабораторных работает внутри конкретного трека. Одинаковые номера в разных треках не считаются дублями, потому что находятся в разных ветках.

## Стандарт папки лабораторной

Каждая полноценная лабораторная хранится в отдельной папке:

- `00_Lab Card.md` — карточка, статус, технологии, связи.
- `01_Task.md` — цель, задание, требования, ожидаемый результат.
- `02_Worklog.md` — ход выполнения, команды, проверки.
- `03_Code.md` — ключевой код или разбор уже написанного кода.
- `04_Errors.md` — ошибки, причины, решения.
- `05_Result.md` — итог, навыки, дальнейшие улучшения.
