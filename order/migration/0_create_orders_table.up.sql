CREATE TABLE IF NOT EXISTS orders(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(32),
    customer_name VARCHAR(64),
    address_line1 VARCHAR(64),
    address_line2 VARCHAR(64),
    address_state VARCHAR(2),
    address_zip VARCHAR(5),
    payment_ccnumber VARCHAR(16),
    payment_ccname VARCHAR(32),
    payment_ccexpires VARCHAR(5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);