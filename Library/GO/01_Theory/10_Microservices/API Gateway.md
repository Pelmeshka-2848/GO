---
type: theory
area: theory
status: active
created: 2026-06-06
updated: 2026-06-09
tags:
  - area/theory
  - microservices
---

# API Gateway

## Кратко

API Gateway — входная точка перед набором сервисов, которая маршрутизирует запросы и применяет общую инфраструктурную политику.

## Зачем нужно

- скрыть внутреннюю топологию сервисов;
- централизовать routing;
- выполнять auth/rate limiting/CORS;
- агрегировать ответы;
- упростить внешний API.

## Правила

- Gateway должен быть тонким.
- Бизнес-логика живет в сервисах.
- У downstream-вызовов должны быть timeout-ы.
- Ошибки сервисов нужно преобразовывать аккуратно.
- Gateway должен участвовать в tracing.

## Типичные ошибки

- превратить gateway в новый монолит;
- не настроить timeout к сервисам;
- скрыть реальные ошибки так, что их нельзя диагностировать;
- сделать gateway единственной точкой авторизации без проверки в сервисах, где она нужна.

## Связанные темы

- [[01_Theory/08_Architecture/API_Gateway]]
- [[01_Theory/10_Microservices/Service Discovery]]
- [[01_Theory/10_Microservices/Distributed Tracing]]

