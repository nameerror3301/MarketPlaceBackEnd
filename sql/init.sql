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



