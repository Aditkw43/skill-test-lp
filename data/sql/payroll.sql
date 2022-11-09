CREATE DATABASE IF NOT EXISTS lp_db;
USE lp_db;
CREATE TABLE IF NOT EXISTS payroll (
    batch INT NOT NULL,
    user_id INT NOT NULL,
    account_name VARCHAR(max) NOT NULL,
    account_number VARCHAR(max) NOT NULL,
    amount INT NOT NULL,
    status TINYINT(1) NOT NULL,
);