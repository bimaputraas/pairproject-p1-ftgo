package handler

import (
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserInfo struct {
	ID int
}

type UserHandler struct {
	DB *sql.DB
}

func (db UserHandler) GetMenu(customerID int) {
	var (
		id      int
		name    string
		price   float64
		alcohol bool
	)
	rows, err := db.DB.Query(`
	SELECT * FROM Beverages
	`)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		rows.Scan(&id, &name, &price, &alcohol)
		if alcohol {

		}
		fmt.Printf("ID: %d, Name: %s, Price: %f, \n", id, name, price)
	}
}

func (db UserHandler) Register(email, password string) error {
	rows, err := db.DB.Query(`
	SELECT 1 FROM Customers WHERE email = ?;
	`, email)
	if err != nil {
		return err
	}
	if rows.Next() {
		return errors.New("Email already exist")
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return err
	}
	_, err = db.DB.Exec(`
	INSERT INTO Customers (email, password)
	VALUES (?, ?);
	`, email, string(hashedPass))
	if err != nil {
		return err
	}
	// Get ID
	var customerID int
	row := db.DB.QueryRow(`
	SELECT id FROM Customer WHERE email = ?
	`, email)
	err = row.Scan(&customerID)
	if err != nil {
		return err
	}
	db.DB.Exec(`
	INSERT INTO CustomerDetails (name, age, phone, customerID)
	VALUES ("", null, null, ?);
	`, customerID)
	return nil
}

func (db UserHandler) Login(email, password string) (int, error) {
	var (
		id         int
		hashedPass []byte
	)
	row := db.DB.QueryRow(`
	SELECT id, password FROM Customer WHERE email = ?
	`, email)
	err := row.Scan(&id, &hashedPass)
	if err != nil {
		return -1, err
	}
	err = bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (db UserHandler) CreateDetails(name string, age int, phone int, customerID int) error {
	rows, err := db.DB.Query(`
	SELECT 1 FROM Customers WHERE id = ?;
	`, customerID)
	if err != nil {
		return err
	}
	if !rows.Next() {
		return errors.New("Customer doesn't exist!")
	}
	db.DB.Exec(`
	INSERT INTO CustomerDetails (name, age, phone, customerID)
	VALUES (?, ?, ?, ?);
	`, name, age, phone, customerID)

	return nil
}

func (db UserHandler) UpdateDetails() {
	// update
}

func (db UserHandler) DeleteDetails() {
	// delete
}
