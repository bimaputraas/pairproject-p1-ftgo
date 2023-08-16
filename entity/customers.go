package entity

type Customers struct {
	Id       int
	Email    string
	Password string
}

// id INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
// email VARCHAR(255) UNIQUE,
// password VARCHAR(255)