CREATE DATABASE example;

\c example;

CREATE TABLE IF NOT EXISTS users (
    id int generated always as identity primary key,
    name varchar(256),
    age int 
);

INSERT INTO users (name, age) VALUES ('tarou', 15), ('hanako', 33);
