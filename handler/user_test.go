package handler

import (
	"fmt"
	"testing"

	"github.com/bimaxputra/pairproject-p1-ftgo/config"
)

func TestLogin(t *testing.T) {
	db := config.ConnectDb()
	defer db.Close()
	uh := Handler{UserHandler: db}
	err := uh.LoginUser("test@mail.com", "test")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Successfully logged in")
}

func TestLoginFail(t *testing.T) {
	db := config.ConnectDb()
	defer db.Close()
	uh := Handler{UserHandler: db}
	err := uh.LoginUser("errorEmail", "errorPass")
	if err == nil {
		t.Error("Logged in into non existant user?")
	}
}
