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
func (h *Handler) RegisterUser(email, password string) string {
	ctx := context.Background()
	query := `SELECT email FROM customers WHERE email = ?;`

	rows, err := h.UserHandler.QueryContext(ctx, query, email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	if rows.Next() {
		fmt.Println("Email already exists")
		return email
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	query2 := `INSERT INTO customers(email,password) VALUES (?,?)`

	_, err = h.UserHandler.ExecContext(ctx, query2, email, string(hashedPassword))
	if err != nil {
		log.Fatal(err)
	}
	return email
}

// login
func (h *Handler) LoginUser(email, plainPassword string) error {
	ctx := context.Background()
	query := `SELECT email,password FROM customers WHERE email = ?;`
	rows, err := h.UserHandler.QueryContext(ctx, query, email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		// get hashed password in database
		var customer entity.Customers
		rows.Scan(&customer.Email, &customer.Password)

		// comparing hashed password in database with hashed password by user input
		err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(plainPassword))
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

	rows, err := h.UserHandler.QueryContext(ctx, query, email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var customer_id int
	if rows.Next() {
		err := rows.Scan(&customer_id)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Selected customer ID: ", customer_id)
	}
	return customer_id
}

// insert customers details by default
func (h *Handler) InsertDefaultCustomersDetails(customer_id int) {
	ctx := context.Background()
	query := `INSERT INTO customers_details (customer_id) VALUES (?);`

	_, err := h.UserHandler.ExecContext(ctx, query, customer_id)
	if err != nil {
		log.Fatal(err)
	}
	// debug
	fmt.Println("Customer's detail inserted successfully")
}

// AFTER LOGIN

// update customers details
func (h *Handler) UpdateCustomersDetails(name string, age int, phone string, customer_id int) {
	ctx := context.Background()
	query := `UPDATE customers_details SET name = ?, age = ?, phone = ? WHERE customer_id = ?;`

	_, err := h.UserHandler.ExecContext(ctx, query, name, age, phone, customer_id)
	if err != nil {
		log.Fatal(err)
	}
	// debug
	fmt.Println("Customer's detail updated successfully")
}

// view customers details
func (h *Handler) ViewCustomersDetails(customer_id int) (*entity.Customers_details, error) {
	ctx := context.Background()
	query := `SELECT name,age,phone FROM customers_details WHERE customer_id = ?`

	rows, err := h.UserHandler.QueryContext(ctx, query, customer_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var customer_details entity.Customers_details
	if rows.Next() {
		err := rows.Scan(&customer_details.Name, &customer_details.Age, &customer_details.Phone)
		if err != nil {
			return nil, errors.New("Data is not available, please update your data")
		}
		// debug
		fmt.Println("Customers info success")

	}
	return &customer_details, nil

	// debug
}

// view menu
func (h *Handler) ViewBeverages() (map[int]entity.Beverages, error) {
	ctx := context.Background()
	query := `SELECT id,name,price,contains_alcohol FROM beverages`
	rows, err := h.UserHandler.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	beverages := make(map[int]entity.Beverages)
	for rows.Next() {
		var bev entity.Beverages
		rows.Scan(&bev.Id, &bev.Name, &bev.Price, &bev.Alcohol)
		beverages[bev.Id] = bev
	}
	return beverages, nil
}
