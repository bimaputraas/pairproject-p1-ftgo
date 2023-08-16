-- ddl

CREATE TABLE customers(
id INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
email VARCHAR(255) UNIQUE,
password VARCHAR(255)
);

CREATE TABLE customers_details(
id INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
name VARCHAR(255) DEFAULT NULL,
age INT DEFAULT NULL,
phone VARCHAR(255) DEFAULT NULL,
customer_id INT UNIQUE,
FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE beverages(
id INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
name VARCHAR(255),
price DECIMAL(10,2),
contains_alcohol BOOLEAN
);

CREATE TABLE orders(
id INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
customer_id INT,
FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE orders_details(
id INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
order_id INT,
beverage_id INT,
quantity INT,
FOREIGN KEY (order_id) REFERENCES orders(id),
FOREIGN KEY (beverage_id) REFERENCES beverages(id)
);

-- dml
INSERT INTO beverages (name, price,contains_alcohol)
VALUES
('Coca-cola', 25.99,false),
('Sprite', 30.99,false),
('Fanta', 27.99,false),
('Mineral Water', 20.99,false),
('Green Sands', 26.99,false),
('Bintang', 40.99,true),
('Anker', 45.99,true),
('Prost', 55.99,true),
('Wine', 99.99,true),
('Soju', 75.99,true);
