# wb-test-exercise
# Version: 1.0
Тестовое задание
Сервис для управления отношениями Складов к Телеграм чатам
#
Шаги развёртки приложения:
1. Создайте .env файл и внесите в него параметр MONGO_DB_PASSWORD - пароль от юзера БД MongoDB
2. Внесите правки в файле `config/app.yml` в блоке `db.mongo`
    - `host` - адрес БД;
    - `port` - порт;
    - `database` - имя БД;
    - `collection` - имя коллекции;
3. `docker build -t office-tg-manager:1.0 .`
4. `docker run -d -p 80:8000 {IMAGE ID}`

# Features
1. GET `/api/{ver}/telegram-to-office-relations/items` - вернуть все записи
2. GET `/api/{ver}/telegram-to-office-relations/items/{objectId}` - вернуть конкретную запись
3. POST `/api/{ver}/telegram-to-office-relations/create` - создать новую запись
4. POST `/api/{ver}/telegram-to-office-relations/update/{objectId}` - изменить существующую запись
5. POST `/api/{ver}/telegram-to-office-relations/delete/{objectId}` - удалить существующую запись
6. GET `/api/metrics` - метрика Prometheus
