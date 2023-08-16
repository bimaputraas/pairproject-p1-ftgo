package entity

type Customers_details struct {
	Id          int
	Name        string
	Age         int
	Phone       string
	Customer_id int
}

// CREATE TABLE customers_details(
// 	id INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
// 	name VARCHAR(255) DEFAULT NULL,
// 	age INT DEFAULT NULL,
// 	phone VARCHAR(255) DEFAULT NULL,
// 	customer_id INT UNIQUE,
// 	FOREIGN KEY (customer_id) REFERENCES customers(id)
// );