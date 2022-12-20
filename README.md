# service for work with currency


## Заполнение конифгов
1. Переименовать файл .env-example на конфиги, которые вы хотите использовать
2. Создать api key для получения текущих валют с сайта https://currate.ru
3. Заполнить им поле CURRENCY_API_KEY

## Запуск сервиса локально:
1. Запуск сервера
```shell
go run main.go
```
2. Миграции в базу данных 
```shell
make migrate-up
```

## Запуск сервиса через Docker:
1. Запуск Docker-compose
```shell
docker compose up
```
2. Миграции в базу данных 
```shell
make migrate-up
```

## Примеры запроса на сервер
POST /api/currency \
Request:
```json
{
    "currencyFrom": "RUB",
    "currencyTo": "USD"
}
```
Response:
```json
{
  "status": "success"
}
```

PUT /api/currency \
Request:
```json
{
  "currencyFrom": "RUB",
  "currencyTo": "USD",
  "value": 99
}
```
Response:
```json
{
  "status": "success"
}
```

GET /api/currency \
Response:
```json
[
  {
    "currencyFrom": "RUB",
    "currencyTo": "USD"
  }
]
```