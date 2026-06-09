# Lab 005 - PostgreSQL CRUD — Errors

| Ошибка | Причина | Решение | Связанная заметка |
|---|---|---|---|
| `connection refused` | PostgreSQL не запущен или указан неверный порт | Проверить службу PostgreSQL, host и port в DSN | [[02_Practice/03_Debug_Diary/PostgreSQL Connection Errors]] |
| `password authentication failed` | Неверный пользователь или пароль | Проверить пользователя, пароль и настройки доступа | [[01_Theory/06_Databases/Connection_String]] |
| `relation does not exist` | Таблица не создана или миграция не применена | Выполнить миграции перед запуском приложения | [[01_Theory/06_Databases/Migrations]] |
