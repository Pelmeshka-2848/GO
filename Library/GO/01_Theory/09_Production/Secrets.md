# Secrets

## Что это

Secret — чувствительное значение, которое нельзя раскрывать: пароль, token, private key, connection string с паролем.

## Зачем нужно

Утечка секрета может дать доступ к базе данных, API, пользователям или инфраструктуре.

## Где хранить

- env-переменные;
- secret manager;
- protected CI/CD variables;
- локальный `.env`, который не попадает в Git.

## Где не хранить

- Go-файлы;
- README;
- Obsidian-заметки с реальными значениями;
- логи;
- screenshots;
- публичные issues.

## Правила

- Если secret попал в Git, считай его скомпрометированным.
- После утечки secret нужно ротировать.
- Логи должны маскировать секреты.
- Разные окружения используют разные секреты.
- Минимизируй доступ к production-секретам.

## Частые ошибки

- закоммитить `.env`;
- вставить token в пример команды;
- отправить connection string в лог;
- использовать один secret для dev и production.

## Связанные темы

- [[01_Theory/90_Rules/Logging Rules]]
- [[01_Theory/09_Production/Config_Env]]
- [[01_Theory/05_HTTP_Backend/01_HTTP_Fundamentals/API_Tokens]]

