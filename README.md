# EWallet

Тестовое задание
_____
Для запуска нужно применить миграции:
```
    migrate -path ./schema -database "postgres://postgres:1488@localhost:5438/postgres?sslmode=disable" up
```
Далее запустить контейнер:
```
    docker-compose up --build ewallet
```
