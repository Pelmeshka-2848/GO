# Vault Structure

Карта структуры Go-хранилища. Используй ее как правило размещения заметок и как быстрый способ понять, куда относить новый материал.

## Главный принцип

Хранилище разделено по роли заметки:

- `00_*` — входы, курс, дорожные карты и временный входящий поток.
- `01_Theory` — чистая теория: понятия, правила, синтаксис, паттерны, справочные объяснения.
- `02_Practice` — практическая работа: лабораторные, кейсы, ошибки, рецепты и заметки проектов.
- `03_Reference` — внешние ссылки, глоссарий и настройки справочного характера.
- `99_Templates` — шаблоны, которые не должны участвовать в рабочем графе как учебные заметки.

## Верхний уровень

| Путь | Назначение | Канонический вход |
|---|---|---|
| `00_Course` | маршрут обучения по модулям | [[00_Go_Course_Map]] |
| `00_Inbox` | сырые мысли и временные заметки | [[00_Inbox/Inbox]] |
| `00_Roadmaps` | дорожные карты, текущий фокус и пробелы | [[00_Roadmaps/Go Learning Roadmap]] |
| `01_Theory` | теория Go и backend-разработки | [[01_Theory/00_Index]] |
| `02_Practice` | лабораторные, кейсы, ошибки, рецепты | [[02_Practice/00_Practice_Index]] |
| `03_Reference` | ссылки, глоссарий, справка | [[03_Reference/Links]] |
| `99_Templates` | шаблоны карточек и заметок | [[99_Templates/00_Templates_Index]] |

## Теория

| Путь | Что хранить |
|---|---|
| `01_Theory/01_Go_Basics` | базовый синтаксис, типы, функции, структуры, интерфейсы |
| `01_Theory/02_Go_Tooling` | `go run`, `go build`, modules, fmt, test, vet, структура проекта |
| `01_Theory/03_CLI_And_Files` | CLI, аргументы, флаги, файлы, JSON, конфиги, строки |
| `01_Theory/04_Stdlib` | пакеты стандартной библиотеки |
| `01_Theory/05_HTTP_Backend` | HTTP, REST, handlers, middleware, JSON API, tokens |
| `01_Theory/06_Databases` | PostgreSQL, SQL, migrations, transactions, repository layer |
| `01_Theory/07_Concurrency` | goroutines, channels, context, mutex, worker pool |
| `01_Theory/08_Architecture` | слои, DTO, DI, Clean Architecture, microservices basics |
| `01_Theory/09_Production` | config, secrets, logging, healthcheck, JWT, observability |
| `01_Theory/10_Microservices` | gateway, discovery, broker, tracing, gRPC |
| `01_Theory/11_Advanced_Go` | testing, benchmarks, pprof, reflection, generics |
| `01_Theory/90_Rules` | правила проектирования и оформления |
| `01_Theory/91_Lifehacks` | практические памятки и приемы мышления |
| `01_Theory/99_Atomic_Go` | маленькие атомарные заметки по конкретному приему |
| `01_Theory/99_Aliases` | старые входные точки и переходники |

## Практика

| Путь | Что хранить |
|---|---|
| `02_Practice/01_Labs` | полноценные лабораторные с карточкой, заданием, worklog, кодом, ошибками и итогом |
| `02_Practice/02_Mini_Cases` | короткие самостоятельные задачи и тренировочные кейсы |
| `02_Practice/02_Development_Cases` | прикладные backend-кейсы: API, CRUD, auth, architecture |
| `02_Practice/03_Debug_Diary` | повторяемые ошибки, причины, решения и проверки |
| `02_Practice/04_Code_Recipes` | универсальные рецепты кода и разборы переиспользуемых фрагментов |
| `02_Practice/05_Project_Notes` | заметки по своим проектам |

## Как выбрать папку

| Вопрос | Куда класть |
|---|---|
| "Как работает концепция?" | `01_Theory` |
| "Как я сделал это в конкретной работе?" | `02_Practice/01_Labs` или `02_Practice/02_Mini_Cases` |
| "Какая ошибка была и как я ее исправил?" | `02_Practice/03_Debug_Diary` |
| "Этот фрагмент можно переиспользовать?" | `02_Practice/04_Code_Recipes` |
| "Это ссылка, термин или справка?" | `03_Reference` |
| "Это сырая мысль без места?" | `00_Inbox` |

## Канонические индексы

- [[Home]]
- [[00_Start_Here]]
- [[00_Go_Course_Map]]
- [[01_Theory/00_Index]]
- [[02_Practice/00_Practice_Index]]
- [[02_Practice/01_Labs/00_Labs_Index]]
- [[02_Practice/02_Mini_Cases/00_Mini_Cases_Index]]
- [[02_Practice/03_Debug_Diary/00_Debug_Diary_Index]]
- [[02_Practice/04_Code_Recipes/00_Code_Recipes_Index]]

## Правила чистоты

- Новые основные индексы называются по схеме `00_*_Index.md`.
- Старые индексы с человекочитаемыми именами можно оставлять как переходники.
- Теория может ссылаться на теорию.
- Практика может ссылаться на теорию, ошибки и рецепты.
- Теория не должна массово ссылаться на каждую лабораторную.
- Шаблоны не связываются с рабочими заметками.
- Для самих шаблонов есть отдельный индекс: [[99_Templates/00_Templates_Index]].
