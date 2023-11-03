-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255),
    password VARCHAR(255),
    deposit_amount NUMERIC
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    stock INT,
    price NUMERIC
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    user_id INT,
    product_id INT,
    store_id INT, 
    quantity INT,
    total_amount NUMERIC
);

