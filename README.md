# test-SkillsRock

REST API для управления задачами с использованием:
- Go 1.23.4
- Fiber (веб-фреймворк)
- PostgreSQL (драйвер pgx)
- Docker

## Запуск проекта

### Требования
- Docker и Docker Compose
- Go 1.23.4+ (если запуск без Docker)

### Инструкция
1. Склонируйте репозиторий:
   ```bash
   git clone https://github.com/ваш-репозиторий.git
   cd test-SkillsRock

2. Запустите сервисы:
    ```bash
    docker-compose up --build
3. API будет доступно на:
    http://localhost:3000


### Эндпоинты

GET /tasks — список всех задач
POST /tasks — создать задачу (JSON: title, description, status)
PUT /tasks/:id — обновить задачу
DELETE /tasks/:id — удалить задачу