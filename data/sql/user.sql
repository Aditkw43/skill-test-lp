CREATE DATABASE IF NOT EXISTS lp_db;
USE lp_db;
CREATE TABLE IF NOT EXISTS user (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name VARCHAR(max) NOT NULL,
    status TINYINT(1) NOT NULL,
);