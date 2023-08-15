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

func (db UserHandler) GetMenu() {
	var (
		id    int
		name  string
		price float64
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
	hashedPass, err:= bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return err
	}
	_, err = db.DB.Exec(`
	INSERT INTO Customers (email, password)
	VALUES (?, ?);
	`, email, hashedPass)
	if err != nil {
		return err
	}
	return nil
}

func (db UserHandler) Login(email, password string) UserInfo, error {
	var hashedPass []byte
	row := db.DB.QueryRow(`
	SELECT password FROM Customer WHERE email = ?
	`, email)
	err := row.Scan(&hashedPass)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (db UserHandler) CreateDetails() {
	// rows, err := db.DB.Query(`
	// SELECT 1 FROM Customers WHERE email = ?;
	// `, email)
	// if err != nil {
	// 	return err
	// }
	if rows.Next() {
		return errors.New("Email already exist")
	}
	db.DB.Exec(`
	INSERT INTO Customer (email, password)
	VALUES (?, ?);
	`)

	return nil
}

func (db UserHandler) UpdateDetails() {

}

func (db UserHandler) DeleteDetails() {

}
