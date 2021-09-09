CREATE DATABASE IF NOT EXISTS photo_gallery;
USE photo_gallery;

CREATE TABLE users IF NOT EXISTS (
    id INT PRIMARY KEY,
    name VARCHAR(30), 
    email VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    id INT PRIMARY KEY,
    user_id INT NOT NULL,
    amount FLOAT(6,2),
    description VARCHAR(255)
);


