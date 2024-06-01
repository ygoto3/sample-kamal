CREATE DATABASE sample_db;

USE sample_db;
 
CREATE TABLE users(
  id int,
  name varchar(50),
  age int
);
 
INSERT INTO users (id, name, age) VALUES (0, 'Mike', 25);
INSERT INTO users (id, name, age) VALUES (1, 'Emma', 23);
 
SELECT *
FROM users;
