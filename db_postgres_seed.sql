-- Crear base de datos
CREATE DATABASE banking;

-- Conectarse a la base de datos
\c banking;

-- Eliminar tablas si existen
DROP TABLE IF EXISTS transactions CASCADE;
DROP TABLE IF EXISTS accounts CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS refresh_token_store CASCADE;
DROP TABLE IF EXISTS customers CASCADE;

-- Crear tabla customers
-- CREATE TABLE customers (
--     customer_id SERIAL PRIMARY KEY,
--     name VARCHAR(100) NOT NULL,
--     date_of_birth DATE NOT NULL,
--     city VARCHAR(100) NOT NULL,
--     zipcode VARCHAR(10) NOT NULL,
--     status BOOLEAN NOT NULL DEFAULT TRUE
-- );

-- -- Insertar datos en customers
-- INSERT INTO customers (customer_id, name, date_of_birth, city, zipcode, status) VALUES
--     (2000, 'Steve', '1978-12-15', 'Delhi', '110075', TRUE),
--     (2001, 'Arian', '1988-05-21', 'Newburgh, NY', '12550', TRUE),
--     (2002, 'Hadley', '1988-04-30', 'Englewood, NJ', '07631', TRUE),
--     (2003, 'Ben', '1988-01-04', 'Manchester, NH', '03102', FALSE),
--     (2004, 'Nina', '1988-05-14', 'Clarkston, MI', '48348', TRUE),
--     (2005, 'Osman', '1988-11-08', 'Hyattsville, MD', '20782', FALSE);

-- Crear la tabla customers con los campos password y role
CREATE TABLE customers (
    customer_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    date_of_birth DATE NOT NULL,
    city VARCHAR(100) NOT NULL,
    zipcode VARCHAR(10) NOT NULL,
    status BOOLEAN NOT NULL DEFAULT TRUE,
    password VARCHAR(255) NOT NULL
    role VARCHAR(50) NOT NULL DEFAULT 'customer'
);

-- Insertar datos en customers con los valores de password y role
INSERT INTO customers (customer_id, name, date_of_birth, city, zipcode, status, password, role) VALUES
    (2000, 'Steve', '1978-12-15', 'Delhi', '110075', TRUE, 'Abc#123', 'customer'),
    (2001, 'Arian', '1988-05-21', 'Newburgh, NY', '12550', TRUE, 'Abc#123', 'customer'),
    (2002, 'Hadley', '1988-04-30', 'Englewood, NJ', '07631', TRUE, 'Abc#123', 'customer'),
    (2003, 'Ben', '1988-01-04', 'Manchester, NH', '03102', FALSE, 'Abc#123', 'customer'),
    (2004, 'Nina', '1988-05-14', 'Clarkston, MI', '48348', TRUE, 'Abc#123', 'customer'),
    (2005, 'Osman', '1988-11-08', 'Hyattsville, MD', '20782', FALSE, 'Abc#123', 'customer');


-- Crear tabla accounts
CREATE TABLE accounts (
    account_id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL REFERENCES customers(customer_id),
    opening_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    account_type VARCHAR(10) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    status BOOLEAN NOT NULL DEFAULT TRUE
);

-- Insertar datos en accounts
INSERT INTO accounts (account_id, customer_id, opening_date, account_type, amount, status) VALUES
    (95470, 2000, '2020-08-22 10:20:06', 'saving', 6823.23, TRUE),
    (95471, 2002, '2020-08-09 10:27:22', 'checking', 3342.96, TRUE),
    (95472, 2001, '2020-08-09 10:35:22', 'saving', 7000, TRUE),
    (95473, 2001, '2020-08-09 10:38:22', 'saving', 5861.86, TRUE);

-- Crear tabla transactions
CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    account_id INT NOT NULL REFERENCES accounts(account_id),
    amount DECIMAL(10,2) NOT NULL,
    transaction_type VARCHAR(10) NOT NULL,
    transaction_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- -- Crear tabla users
-- CREATE TABLE users (
--     username VARCHAR(20) PRIMARY KEY,
--     password VARCHAR(20) NOT NULL,
--     role VARCHAR(20) NOT NULL,
--     customer_id INT REFERENCES customers(customer_id),
--     created_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
-- );

-- -- Insertar datos en users
-- INSERT INTO users (username, password, role, customer_id, created_on) VALUES
--     ('admin', 'abc123', 'admin', NULL, '2020-08-09 10:27:22'),
--     ('2001', 'abc123', 'user', 2001, '2020-08-09 10:27:22'),
--     ('2000', 'abc123', 'user', 2000, '2020-08-09 10:27:22');

-- Crear la tabla users con user_id como identificador único y username como el nombre del usuario
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,  -- user_id como identificador único
    username VARCHAR(100) NOT NULL,  -- username será el nombre del usuario
    password VARCHAR(255) NOT NULL,  -- Cambié el tamaño de password por seguridad
    role VARCHAR(20) NOT NULL,
    customer_id INT REFERENCES customers(customer_id),
    created_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Insertar datos en users con username como el nombre del usuario
INSERT INTO users (username, password, role, customer_id, created_on) VALUES
    ('admin_user', 'abc123', 'admin', NULL, '2020-08-09 10:27:22'),  -- admin_user como nombre de usuario
    ('Steve', 'abc123', 'user', 2000, '2020-08-09 10:27:22'),  -- 'Steve' como nombre de usuario
    ('Arian', 'abc123', 'user', 2001, '2020-08-09 10:27:22');  -- 'Arian' como nombre de usuario


-- Crear tabla refresh_token_store
CREATE TABLE refresh_token_store (
    refresh_token VARCHAR(300) PRIMARY KEY,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
