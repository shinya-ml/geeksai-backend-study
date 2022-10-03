CREATE DATABASE example;

\c example;

CREATE SCHEMA example;

CREATE TABLE example.user (
    id int generated always as identity primary key,
    name varchar(256),
    age int 
);

INSERT INTO example.user (name, age) VALUES ('tarou', 15), ('hanako', 33);
