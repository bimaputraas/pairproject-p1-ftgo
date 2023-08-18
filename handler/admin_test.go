package handler

import (
	"testing"

	"github.com/bimaxputra/pairproject-p1-ftgo/config"
)

func TestDelete(t *testing.T) {
	db := config.ConnectDb()
	defer db.Close()
	uh := Handler{UserHandler: db}
	err := uh.AddBeverage("test", 1, true)
	if err != nil {
		t.Error("Failed on adding beverage")
	}

	var id int
	row := uh.UserHandler.QueryRow(`SELECT id FROM beverages WHERE name = "test"`)
	err = row.Scan(&id)
	if err != nil {
		t.Error("Failed on finding id")
	}
	err = uh.DeleteBeveragebyId(id)
	if err != nil {
		t.Error("Failed on deleting beverage")
	}
}

func TestDeleteFail(t *testing.T) {
	db := config.ConnectDb()
	defer db.Close()
	uh := Handler{UserHandler: db}
	err := uh.AddBeverage("test", 1, true)
	if err != nil {
		t.Error("Failed on adding beverage")
	}

	err = uh.DeleteBeveragebyId(1)
	if err == nil {
		t.Error("id=1 suppose to have others depending on it")
	}
}
