# 🌍 GeoService Controller

![GeoService Banner](geoservice_banner_1777367185368.png)

[![Go Version](https://img.shields.io/badge/Go-1.24-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=flat-square)](LICENSE)
[![Docker](https://img.shields.io/badge/Docker-Supported-2496ED?style=flat-square&logo=docker)](https://www.docker.com/)
[![Prometheus](https://img.shields.io/badge/Monitoring-Prometheus-e6441c?style=flat-square&logo=prometheus)](https://prometheus.io/)

**GeoService Controller** — это высокопроизводительный микросервис на Go для геокодирования и поиска адресов, интегрированный с DaData API. Сервис включает в себя полноценную систему аутентификации, многоуровневое кэширование и продвинутый мониторинг.

---

## ✨ Основные возможности

- 🔍 **Поиск адресов** — интеллектуальный поиск по текстовому запросу.
- 📍 **Обратное геокодирование** — получение адреса по координатам (lat/lng).
- 🔐 **Безопасность** — JWT аутентификация с использованием HTTP-only Cookies.
- ⚡ **Производительность** — кэширование ответов DaData в Redis для мгновенного доступа.
- 📊 **Наблюдаемость** — детальные метрики (Prometheus) и готовые дашборды (Grafana).
- 📖 **Документация** — интерактивный Swagger UI для быстрого тестирования API.
- 🐳 **Контейнеризация** — легкое развертывание через Docker и Docker Compose.

---

## 🛠 Технологический стек

<p align="left">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
  <img src="https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" />
  <img src="https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white" />
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" />
  <img src="https://img.shields.io/badge/Prometheus-E6441C?style=for-the-badge&logo=prometheus&logoColor=white" />
  <img src="https://img.shields.io/badge/Grafana-F46800?style=for-the-badge&logo=grafana&logoColor=white" />
</p>

- **Core:** Go 1.24, Chi Router
- **Database:** PostgreSQL (pgx), Redis (go-redis)
- **Security:** JWT Auth (jwtauth), Argon2/bcrypt
- **Metrics:** Prometheus Client, Grafana
- **API Docs:** Swag (Swagger 2.0)

---

## 🚀 Быстрый старт

### Требования
- Docker & Docker Compose
- DaData API Key & Secret ([получить здесь](https://dadata.ru/))

### Установка

1. **Клонируйте репозиторий:**
   ```bash
   git clone https://github.com/yourusername/geoservise-controller.git
   cd geoservise-controller
   ```

2. **Настройте окружение:**
   Создайте файл `.env` на основе примера:
   ```bash
   cp .env.example .env
   ```
   *Обязательно заполните `DADATA_API_KEY` и `DADATA_SECRET_KEY`.*

3. **Запустите проект:**
   ```bash
   docker compose up --build -d
   ```

---

## 📂 Структура проекта

```text
├── cmd/geo/              # Точка входа в приложение
├── internal/
│   ├── app/              # Сборка и инициализация зависимостей
│   ├── delivery/         # HTTP слой (обработчики, middleware)
│   ├── infrastructure/   # Внешние сервисы (БД, Redis, Auth)
│   ├── model/            # Доменные модели данных
│   └── usecase/          # Бизнес-логика приложения
├── docs/                 # Автогенерируемая документация Swagger
├── grafana/              # Конфигурация и дашборды мониторинга
└── compose.yaml          # Описание инфраструктуры
```

---

## 📡 API Эндпоинты

### Аутентификация
- `POST /api/register` — Регистрация нового аккаунта.
- `POST /api/login` — Вход (установка JWT Cookie).

### Геосервисы (защищены JWT)
- `POST /api/address/search` — Поиск по строке.
- `POST /api/address/geocode` — Координаты → Адрес.

### Сервисные
- `GET /swagger/index.html` — Документация API.
- `GET /metrics` — Метрики для Prometheus.

---

## 📊 Мониторинг и метрики

Сервис предоставляет глубокую аналитику в реальном времени:
- **Prometheus:** Доступен на порту `9090`.
- **Grafana:** Доступна на порту `3000` (логин: `admin`, пароль: `admin`).
- **Следим за:** временем ответа БД, кэш-хитами Redis, статусом DaData API и RPS.

---

## 🤝 Контрибьютинг

Буду рад вашим Pull Requests! Для серьезных изменений, пожалуйста, сначала создайте Issue, чтобы обсудить, что вы хотите изменить.

---

## 📜 Лицензия

Проект распространяется под лицензией MIT. Подробности в файле [LICENSE](LICENSE).

---

<p align="center">
  Сделано с ❤️ для разработчиков.
</p>
