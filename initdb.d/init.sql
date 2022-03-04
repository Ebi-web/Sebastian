DROP
DATABASE IF EXISTS sebastian;
CREATE
DATABASE sebastian;
USE sebastian;

DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id         bigint PRIMARY KEY,
    name       text      NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin