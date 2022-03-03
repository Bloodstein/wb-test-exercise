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