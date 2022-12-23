# MarketPlaceBackEnd

## Описание

Проект был реализован на комерчесской основе. Требования были просты, нужен был API для маркет плейса по продаже товаров.

К сожалению по требованию со стороны заказчика, предоставить технические требования не могу.

## Структура проекта

```
├── cmd
│   └── app
│       └── main.go
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   └── run.go
│   ├── config
│   │   ├── config.go
│   │   └── config.yaml
│   ├── database
│   │   └── database.go
│   ├── middleware
│   │   ├── content.go
│   │   └── jwt.go
│   ├── models
│   │   ├── products.go
│   │   └── users.go
│   └── routes
│       ├── auth.go
│       ├── products.go
│       └── route.go
├── License.md
├── Readme.md
├── sql
│   └── init.sql
└── tools
    └── crypt.go
```

## API точки доступа

Ниже будут перечисленны точки доступа к API.

URL - http://localhost:8080 

#### **/api/v1/signUp**
* `POST` : Регистрация.   

Пример запроса:

```
{
    "email": "example@gmail.com", 
    "pass": "qwerty"
}
```

Привер ответа:

```
{
    "api_version": "1.0",
    "content": null,
    "description": "Registration was successful!",
    "status_code": 200
}
```

#### **/api/v1/signIn**
* `POST` : Получение Bearer токена. Срок его действия 12 минут.

Пример запроса:

```
{
    "email": "example@gmail.com", 
    "pass": "qwerty"
}
```

Привер ответа:

```
{
    "api_version": "1.0",
    "content": null,
    "description": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzE4MDEyNTV9.wcM4YN9BdRBb1kqdAhNKTRm3lPQ2zQuzQoyQNq5ya2A",
    "status_code": 200
}
```


#### **/api/v1/products**
* `GET` : Получение всех товаров. * `AUTH`

Пример ответа:
```
{
    "api_version": "1.0",
    "content": [
        {
            "id": "1",
            "market_name": "ЯндексМаркет",
            "prod_manufacturer": "Жостовская фабрика декоративной росписи",
            "prod_name": "Жостовский поднос",
            "art": "A5D13012022001",
            "price": "7650",
            "link": "https://market.yandex.ru/offer/jfyvHgj69IT2bDuLZUOlgw"
        },
        {
            "id": "2",
            "market_name": "ЯндексМаркет",
            "prod_manufacturer": "Жостовская фабрика декоративной росписи",
            "prod_name": "Жостовский поднос",
            "art": "A14D30082022016",
            "price": "7155",
            "link": "https://market.yandex.ru/offer/1Ty7-X0Cmw49ROMSP6VniA"
        },
        {
            "id": "3",
            "market_name": "ЯндексМаркет",
            "prod_manufacturer": "Жостовская фабрика декоративной росписи",
            "prod_name": "Жостовский поднос",
            "art": "A14D30082022019",
            "price": "7155",
            "link": "https://market.yandex.ru/offer/mU4yXbo9NwZBGINAifuv-w"
        },
    ],
    "description": "By id product",
    "status_code": 200
}
```


#### **/api/v1/products/1**
* `GET` : Получение товара по его ID. * `AUTH`

Пример ответа:
```
{
    "api_version": "1.0",
    "content": [
        {
            "id": "1",
            "market_name": "ЯндексМаркет",
            "prod_manufacturer": "Жостовская фабрика декоративной росписи",
            "prod_name": "Жостовский поднос",
            "art": "A14D30082022014",
            "price": "6290",
            "link": "https://market.yandex.ru/offer/7_pK0iu7nzWZ5kAmXdwnHg"
        }
    ],
    "description": "By id product",
    "status_code": 200
}
```

#### **/api/v1/products?total=3**
* `GET` : Получение определенного кол-ва товаров. * `AUTH`

Пример ответа:
```
{
    "api_version": "1.0",
    "content": [
        {
            "id": "1",
            "market_name": "ЯндексМаркет",
            "prod_manufacturer": "Жостовская фабрика декоративной росписи",
            "prod_name": "Жостовский поднос",
            "art": "A5D13012022001",
            "price": "7650",
            "link": "https://market.yandex.ru/offer/jfyvHgj69IT2bDuLZUOlgw"
        },
        {
            "id": "2",
            "market_name": "ЯндексМаркет",
            "prod_manufacturer": "Жостовская фабрика декоративной росписи",
            "prod_name": "Жостовский поднос",
            "art": "A14D30082022016",
            "price": "7155",
            "link": "https://market.yandex.ru/offer/1Ty7-X0Cmw49ROMSP6VniA"
        },
        {
            "id": "3",
            "market_name": "ЯндексМаркет",
            "prod_manufacturer": "Жостовская фабрика декоративной росписи",
            "prod_name": "Жостовский поднос",
            "art": "A14D30082022019",
            "price": "7155",
            "link": "https://market.yandex.ru/offer/mU4yXbo9NwZBGINAifuv-w"
        },
    ],
    "description": "By id product",
    "status_code": 200
}
```


## Использование

Для запуска можно использовать **Docker** или же **Golang**.

**Golang** - Для запуска в директории **/cmd/app/** запустите файл **main.go**
**Docker** - Для запуска в корневой директории запустите команду ```docker-compose up --build```

## Разработчики

- [OneByteForLife](https://github.com/nameerror3301)

## Лицензия

- Это программное обеспечение защищено лицензией MIT!
