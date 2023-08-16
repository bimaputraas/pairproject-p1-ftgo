package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bimaxputra/pairproject-p1-ftgo/handler"
)

type Cli struct {
	Handler *handler.Handler
}

func (cli *Cli) MainGateInterface() {
	fmt.Println("[COMMAND]			-DESCRIPTION")
	fmt.Println("[login]				-login")
	fmt.Println("[register]			-register")
	fmt.Printf("\nEnter your command : ")

	askInputGate := ScanInputString()

	switch askInputGate{
	case "login":
		cli.LoginInterface()
	case "register":
		cli.RegisterInterface()
	default :
		fmt.Printf("Invalid Input\n\n")
		cli.MainGateInterface()
	}
	
}

func (cli *Cli) RegisterInterface() {
	fmt.Printf("\n-REGISTER-\n\n")
	fmt.Printf("Please input your email address: ")
	askRegisterEmail := ScanInputString()
	fmt.Printf("Please input your password: ")
	askRegisterPassword := ScanInputString()
	
	// register handler.
	RegisteredEmail := cli.Handler.RegisterUser(askRegisterEmail,askRegisterPassword)
	
	// handler, select id customer by their email, return id customer
	customerId := cli.Handler.SelectByEmail(askRegisterEmail)
	
	// handler,insert customerId to customers_details and set default another data
	cli.Handler.InsertDefaultCustomersDetails(customerId)
	
	fmt.Println("Success register using email",RegisteredEmail)
	fmt.Println("")
	cli.MainGateInterface()
}

func (cli *Cli) LoginInterface() {
	fmt.Printf("\n-LOGIN-\n\n")
	fmt.Printf("Please input your email address: ")
	askInputEmail := ScanInputString()
	fmt.Printf("Please input your password: ")
	askInputPassword := ScanInputString()
	err := cli.Handler.LoginUser(askInputEmail,askInputPassword)

	if err != nil {
		// failed and back to main gate interface
		fmt.Printf("Failed to login\n\n")
		cli.MainGateInterface()
	}
	
	// else (login succes)
	// entering main menu interface
	fmt.Println("login succes!")
	customerId := cli.Handler.SelectByEmail(askInputEmail)
	cli.MainMenuInterface(customerId)
}



// after login
func (cli *Cli) MainMenuInterface(customerId int) {
	// emailCustomer = email dari user yg berhasil login
	fmt.Printf("\nWelcome to main menu!\n\n")
	fmt.Println("[COMMAND]			-DESCRIPTION")
	fmt.Println("[profile]			-User Profile")
	fmt.Printf("\nEnter your command : ")
	
	askInputMain := ScanInputString()
	
	switch askInputMain{
	case "profile":
		cli.UserProfileInterface(customerId)
	default :
		fmt.Printf("\nProgram end, thanks!\n\n")
		return
	}
}

func (cli *Cli) UserProfileInterface(customerId int) {
	fmt.Printf("\nUser Profile\n\n")
	fmt.Println("[COMMAND]			-DESCRIPTION")
	fmt.Println("[view]				-View User Profile")
	fmt.Println("[update]			-Update User Profile")
	fmt.Printf("\nEnter your command : ")
	
	askInputUserProfile := ScanInputString()

	switch askInputUserProfile{
	case "view":
		cli.ViewUserProfileInterface(customerId)
	case "update":
		cli.UpdateUserProfileInterface(customerId)
	default:
		cli.MainMenuInterface(customerId)
	}
}

func (cli *Cli) UpdateUserProfileInterface(customerId int) {
	fmt.Printf("\n-UPDATE PROFILE-\n\n")
	// name
	fmt.Printf("\nPlease input your name: ")
	askUpdateName := ScanInputString()
	// age
	fmt.Printf("Please input your age: ")
	askUpdateAgeStr := ScanInputString()
	askUpdateAgeInt,err := strconv.Atoi(askUpdateAgeStr)
	if err != nil {
		log.Fatal(err)
	}
	// phone
	fmt.Printf("Please input your phone number: ")
	askUpdatePhone := ScanInputString()

	cli.Handler.UpdateCustomersDetails(askUpdateName,askUpdateAgeInt,askUpdatePhone,customerId)
	cli.UserProfileInterface(customerId)
}

// read
func (cli *Cli) ViewUserProfileInterface(customerId int) {
	fmt.Printf("\n-CUSTOMER DETAILS-\n\n")
	customer_details,err := cli.Handler.ViewCustomersDetails(customerId)
	if err != nil {
		fmt.Println(err)
		cli.UserProfileInterface(customerId)
	}

	fmt.Printf("\nName: %s\nAge: %d\nPhone: %s\n\n",customer_details.Name,customer_details.Age,customer_details.Phone)
	cli.BackToMainMenu(customerId)
}

func(cli *Cli) BackToMainMenu(customerId int){
	fmt.Println("Press enter to back to main menu : ")
	backInput := ScanInputString()
	fmt.Println("")
	
	switch backInput{
	case "" :
		cli.MainMenuInterface(customerId)
	default:
		cli.BackToMainMenu(customerId)
	}
}