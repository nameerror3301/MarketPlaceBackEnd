CREATE TABLE users
(
    id SERIAL NOT NULL PRIMARY KEY,
    email VARCHAR(100) UNIQUE,
    hash_pass VARCHAR(255)
);

CREATE TABLE products
(
    id SERIAL NOT NULL PRIMARY KEY,
    market_name VARCHAR(100),
    prod_manufacturer VARCHAR(200),
    prod_name VARCHAR(100),
    art VARCHAR(255) UNIQUE,
    price BIGINT,
    link TEXT
);

INSERT INTO products (market_name, prod_manufacturer, prod_name, art, price, link) VALUES 
('ЯндексМаркет', 'Жостовская фабрика декоративной росписи', 'Жостовский поднос', 'A5D13012022001', 7650, 'https://market.yandex.ru/offer/jfyvHgj69IT2bDuLZUOlgw'),
('ЯндексМаркет', 'Жостовская фабрика декоративной росписи', 'Жостовский поднос', 'A14D30082022016', 7155, 'https://market.yandex.ru/offer/1Ty7-X0Cmw49ROMSP6VniA'),
('ЯндексМаркет', 'Жостовская фабрика декоративной росписи', 'Жостовский поднос', 'A14D30082022019', 7155, 'https://market.yandex.ru/offer/mU4yXbo9NwZBGINAifuv-w'),
('ЯндексМаркет', 'Жостовская фабрика декоративной росписи', 'Жостовский поднос', 'A5D11072022004', 2070, 'https://market.yandex.ru/offer/HLANBUuXrNJ-MxwKgKx_Bg'),
('ЯндексМаркет', 'Жостовская фабрика декоративной росписи', 'Жостовский поднос', 'A14D30082022014', 6290, 'https://market.yandex.ru/offer/7_pK0iu7nzWZ5kAmXdwnHg'),
('ЯндексМаркет', 'Жостовская фабрика декоративной росписи', 'Жостовский поднос', 'A1D30082022015', 1760, 'https://market.yandex.ru/offer/jF8ywQhoR7vaLFiNigAMnQ'),
('ЯндексМаркет', 'Жостовская фабрика декоративной росписи', 'Жостовский поднос', 'A1D11072022004', 1500, 'https://market.yandex.ru/offer/WM-UljZZV7X1Y0-uUdPt5g'),
('ЯндексМаркет', 'Жостовская фабрика декоративной росписи', 'Жостовский поднос', 'A15D30082022013', 5500, 'https://market.yandex.ru/offer/wl5D37eieuMjkbcLIQzX3g'),
('ЯндексМаркет', 'Жостовская фабрика декоративной росписи', 'Жостовский поднос', 'A2D22092022001', 1650, 'https://market.yandex.ru/offer/K3bzG-jxgWpp9S8wAc-ujA'),
('ЯндексМаркет', 'Жостовская фабрика декоративной росписи', 'Жостовский поднос', 'A2D30082022024', 1650, 'https://market.yandex.ru/offer/9ILMllLo6du3NoT-I-DTVw');



