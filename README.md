# EWallet

Тестовое задание
_____
Для запуска сначала сбилдим контейнер:
```
docker-compose build ewallet
```
Далее запустим сам контейнер:
```
docker-compose up ewallet
```
Для запуска впервые нужно применить миграции:
```
migrate -path ./schema -database "postgres://postgres:1488@localhost:5438/postgres?sslmode=disable" up
```
