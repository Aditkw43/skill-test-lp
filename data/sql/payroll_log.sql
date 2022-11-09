CREATE DATABASE IF NOT EXISTS lp_db;
USE lp_db;
CREATE TABLE IF NOT EXISTS payroll_log (
    file_name INT NOT NULL,
    batch INT NOT NULL,
    total_success INT NOT NULL,
    total_failed INT NOT NULL,
    created_at INT NOT NULL,
    created_by INT NOT NULL
);