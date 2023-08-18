package handler

import (
	"context"
	"log"
)

// admin feature

func (h *Handler) AddBeverage(name string, price float64, alcohol bool) error {
	ctx := context.Background()
	query := `INSERT INTO beverages (name, price,contains_alcohol) VALUES (?,?,?);`
	_, err := h.UserHandler.ExecContext(ctx, query, name, price, alcohol)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (h *Handler) DeleteBeveragebyId(id int) error {
	ctx := context.Background()
	query := `DELETE FROM orders_details WHERE id = ?;`
	query2 := `DELETE FROM beverages WHERE id = ?;`
	_, err := h.UserHandler.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	_, err = h.UserHandler.ExecContext(ctx, query2, id)
	if err != nil {
		return err
	}

	return nil
}
