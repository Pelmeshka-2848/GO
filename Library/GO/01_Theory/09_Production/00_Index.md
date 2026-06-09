# Production Index

Раздел про базовую готовность Go backend-приложения к эксплуатации: конфигурация, секреты, auth, JWT, логирование, healthcheck, observability, graceful shutdown и OpenAPI.

## Рекомендуемый порядок

1. [[01_Theory/09_Production/Config_Env|Config Env]]
2. [[01_Theory/09_Production/Secrets|Secrets]]
3. [[01_Theory/09_Production/Logging|Logging]]
4. [[01_Theory/09_Production/Healthcheck|Healthcheck]]
5. [[01_Theory/09_Production/Observability|Observability]]
6. [[01_Theory/09_Production/Graceful_Shutdown|Graceful Shutdown]]
7. [[01_Theory/09_Production/Auth|Auth]]
8. [[01_Theory/09_Production/Password_Hashing|Password Hashing]]
9. [[01_Theory/09_Production/JWT|JWT]]
10. [[01_Theory/09_Production/Swagger_OpenAPI|Swagger OpenAPI]]

## Что нужно уметь

- не хранить секреты в коде;
- запускать один и тот же код с разными настройками;
- писать полезные логи;
- проверять состояние сервиса через healthcheck;
- корректно завершать сервер;
- понимать базовую схему auth/JWT;
- документировать REST API.

