package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/bimaxputra/pairproject-p1-ftgo/entity"

	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	UserHandler *sql.DB
}

// register
func (h *Handler) RegisterUser(email,password string) string{
	ctx := context.Background()
	query := `SELECT email FROM customers WHERE email = ?;`

	rows,err := h.UserHandler.QueryContext(ctx,query,email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	if rows.Next(){
		fmt.Println("Email is already exist")
		return email
	}

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)

	query2 := `INSERT INTO customers(email,password) VALUES (?,?)`

	_,err = h.UserHandler.ExecContext(ctx,query2,email,string(hashedPassword))
	if err != nil {
		log.Fatal(err)
	}
	return email
}

// login
func (h *Handler) LoginUser(email,plainPassword string) error {
	ctx := context.Background()
	query := `SELECT email,password FROM customers WHERE email = ?;`
	rows,err := h.UserHandler.QueryContext(ctx,query,email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next(){
		// get hashed password in database
		var customer entity.Customers
		rows.Scan(&customer.Email,&customer.Password)

		// comparing hashed password in database with hashed password by user input
		err := bcrypt.CompareHashAndPassword([]byte(customer.Password),[]byte(plainPassword))
		if err != nil {
			log.Fatal(err)
		}
		// if compare success(err = nil)
		return nil
	}
	// else(failed)
	return errors.New("Failed login")
}

// select Customer id by email
func (h *Handler) SelectByEmail(email string) int {
	ctx := context.Background()
	query := `SELECT id FROM customers WHERE email = ?;`

	rows,err := h.UserHandler.QueryContext(ctx,query,email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var customer_id int
	if rows.Next(){
		err:=rows.Scan(&customer_id)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("selected id customer : ",customer_id)
	}
	return customer_id
}

// insert customers details by default
func (h *Handler) InsertDefaultCustomersDetails(customer_id int) {
	ctx := context.Background()
	query := `INSERT INTO customers_details (customer_id) VALUES (?);`

	_,err := h.UserHandler.ExecContext(ctx,query,customer_id)
	if err != nil {
		log.Fatal(err)
	}
	// debug
	fmt.Println("success insert default customers details")
}


// AFTER LOGIN

// update customers details
func (h *Handler) UpdateCustomersDetails(name string, age int,phone string,customer_id int) {
	ctx := context.Background()
	query := `UPDATE customers_details SET name = ?, age = ?, phone = ? WHERE customer_id = ?;`
	
	_,err := h.UserHandler.ExecContext(ctx,query,name,age,phone,customer_id)
	if err != nil {
		log.Fatal(err)
	}
	// debug
	fmt.Println("Customers detail updated")
}