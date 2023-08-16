package cli

import (
	"fmt"
	"m_p1/handler"
)

type Cli struct {
	Handler *handler.Handler
}

func (cli *Cli) MainGateInterface() {
	fmt.Println("[COMMAND]			-DESCRIPTION")
	fmt.Println("[login]				-login")
	fmt.Println("[register]			-register")
	fmt.Printf("Enter your command : ")

	askInputGate := ScanInputString()

	switch askInputGate{
	case "login":
		cli.LoginInterface()
	case "register":
		cli.RegisterInterface()
	default :
		fmt.Println("Invalid Input")
	}
	
}

func (cli *Cli) LoginInterface() {
	fmt.Println("LOGIN")
	fmt.Printf("Please input your email address: ")
	askInputEmail := ScanInputString()
	fmt.Printf("Please input your password: ")
	askInputPassword := ScanInputString()
	err := cli.Handler.LoginUser(askInputEmail,askInputPassword)

	if err != nil {
		// failed and back to main gate interface
		cli.MainGateInterface()
		fmt.Println("Failed to login")
	}
	
	// else (login succes)
	// entering main menu interface
	fmt.Println("login succes!")
	cli.MainMenuInterface(askInputEmail)
}

func (cli *Cli) RegisterInterface() {
	fmt.Println("REGISTER")
	fmt.Printf("Please input your email address: ")
	askRegisterEmail := ScanInputString()
	fmt.Printf("Please input your password: ")
	askRegisterPassword := ScanInputString()
	
	// register handler.
	RegisteredEmail := cli.Handler.RegisterUser(askRegisterEmail,askRegisterPassword)

	// handler, select id customer by their email, return id customer
	customer_id := cli.Handler.SelectByEmail(askRegisterEmail)

	// handler,insert customer_id to customers_details and set default another data
	cli.Handler.InsertDefaultCustomersDetails(customer_id)

	fmt.Println("Success register using email",RegisteredEmail)
	fmt.Println("")
	cli.MainGateInterface()
}

// after login
func (cli *Cli) MainMenuInterface(emailCustomer string) {
	// emailCustomer = email dari user yg berhasil login
	customer_id := cli.Handler.SelectByEmail(emailCustomer)
	fmt.Println("Welcome to main menu!")
	fmt.Println("[COMMAND]			-DESCRIPTION")
	fmt.Println("[profile]				-User Profile")
	fmt.Printf("Enter your command : ")

	askInputMain := ScanInputString()

	switch askInputMain{
	case "profile":
		cli.UserProfileInterface(customer_id)
		default :
		fmt.Println("Invalid Input")
	}
}

func (cli *Cli) UserProfileInterface(customerId int) {
	fmt.Println("User Profile")
}