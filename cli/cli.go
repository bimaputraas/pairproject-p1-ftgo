package cli

import (
	"context"
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
	fmt.Println("[exit]				-exit")
	fmt.Printf("\nEnter your command : ")

	askInputGate := ScanInputString()

	switch askInputGate {
	case "login":
		cli.LoginInterface()
	case "register":
		cli.RegisterInterface()
	case "exit":
		fmt.Printf("\nProgram End, Thanks!\n\n")
		return
	default:
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
	RegisteredEmail := cli.Handler.RegisterUser(askRegisterEmail, askRegisterPassword)

	// handler, select id customer by their email, return id customer
	customerId := cli.Handler.SelectByEmail(askRegisterEmail)

	// handler,insert customerId to customers_details and set default another data
	cli.Handler.InsertDefaultCustomersDetails(customerId)

	fmt.Println("Success registering using email", RegisteredEmail)
	fmt.Println("")
	cli.MainGateInterface()
}

func (cli *Cli) LoginInterface() {
	fmt.Printf("\n-LOGIN-\n\n")
	fmt.Printf("Please input your email address: ")
	askInputEmail := ScanInputString()
	fmt.Printf("Please input your password: ")
	askInputPassword := ScanInputString()

	// login by admin
	if askInputEmail == "admin" && askInputPassword == "admin"{
		cli.MainMenuAdmin()
	}

	err := cli.Handler.LoginUser(askInputEmail, askInputPassword)
	if err != nil {
		// failed and back to main gate interface
		fmt.Printf("Failed to login\n\n")
		cli.MainGateInterface()
	}

	// else (login success)
	// entering main menu interface
	fmt.Println("login success!")
	customerId := cli.Handler.SelectByEmail(askInputEmail)
	cli.MainMenuInterface(customerId)
}

// after login
func (cli *Cli) MainMenuInterface(customerId int) {
	// emailCustomer = email dari user yg berhasil login
	fmt.Printf("\nWelcome to main menu!\n\n")
	fmt.Println("[COMMAND]			-DESCRIPTION")
	fmt.Println("[profile]			-User Profile")
	fmt.Println("[order]				-Create a new order")
	fmt.Println("[exit]				-exit program")
	fmt.Printf("\nEnter your command : ")

	askInputMain := ScanInputString()

	switch askInputMain {
	case "profile":
		cli.UserProfileInterface(customerId)
	case "order":
		cli.OrderInterface(customerId)
	case "exit":
		fmt.Printf("\nProgram End, Thanks!\n\n")
		return
	default:
		cli.MainMenuInterface(customerId)
	}
}

func (cli *Cli) UserProfileInterface(customerId int) {
	fmt.Printf("\nUser Profile\n\n")
	fmt.Println("[COMMAND]			-DESCRIPTION")
	fmt.Println("[view]				-View User Profile")
	fmt.Println("[update]			-Update User Profile")
	fmt.Printf("\nEnter your command : ")

	askInputUserProfile := ScanInputString()

	switch askInputUserProfile {
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
	askUpdateAgeInt, err := strconv.Atoi(askUpdateAgeStr)
	if err != nil {
		log.Fatal(err)
	}
	// phone
	fmt.Printf("Please input your phone number: ")
	askUpdatePhone := ScanInputString()

	cli.Handler.UpdateCustomersDetails(askUpdateName, askUpdateAgeInt, askUpdatePhone, customerId)
	cli.UserProfileInterface(customerId)
}

// read
func (cli *Cli) ViewUserProfileInterface(customerId int) {
	fmt.Printf("\n-CUSTOMER DETAILS-\n\n")
	customer_details, err := cli.Handler.ViewCustomersDetails(customerId)
	if err != nil {
		fmt.Println(err)
		cli.UserProfileInterface(customerId)
	}

	fmt.Printf("\nName: %s\nAge: %d\nPhone: %s\n\n", customer_details.Name, customer_details.Age, customer_details.Phone)
	cli.BackToMainMenu(customerId)
}

func (cli *Cli) BackToMainMenu(customerId int) {
	fmt.Println("Press enter to back to main menu : ")
	backInput := ScanInputString()
	fmt.Println("")

	switch backInput {
	case "":
		cli.MainMenuInterface(customerId)
	default:
		cli.BackToMainMenu(customerId)
	}
}
func (cli *Cli) OrderInterface(customerId int) {
	ctx := context.Background()
	//create order item and get the id
	_, err := cli.Handler.UserHandler.ExecContext(ctx, "INSERT INTO orders (customer_id) VALUES (?);", customerId)
	if err != nil {
		fmt.Println(err)
		cli.MainMenuInterface(customerId)
	}
	rows, err := cli.Handler.UserHandler.QueryContext(ctx, "SELECT id FROM orders WHERE customer_id = ? ORDER BY id DESC", customerId)
	if err != nil {
		fmt.Println(err)
		cli.MainMenuInterface(customerId)
	}
	rows.Next()
	var orderId int
	rows.Scan(&orderId)

	//init order details query
	ordersDetailsQuery := "INSERT INTO orders_details (order_id, beverage_id, quantity) VALUES"

	var total float64 = 0

	//get menu
	menu, err := cli.Handler.ViewBeverages()
	if err != nil {
		fmt.Println(err)
		cli.MainMenuInterface(customerId)
	}

	//init bool for is customer of age
	customer_details, err := cli.Handler.ViewCustomersDetails(customerId)
	if err != nil {
		fmt.Println(err)
		cli.MainMenuInterface(customerId)
	}
	var isAdult bool = false
	if customer_details.Age >= 21 {
		isAdult = true
	}
	for {
		//show menu
		println("ID  Name                Price")
		for _, bev := range menu {
			// print bev with string padding
			if bev.Alcohol && isAdult {
				fmt.Printf("%-3s %-20s %.2f Alcoholic\n", fmt.Sprint(bev.Id), bev.Name, bev.Price)
			} else {
				fmt.Printf("%-3s %-20s %.2f\n", fmt.Sprint(bev.Id), bev.Name, bev.Price)
			}
		}

		//get order with cli
		fmt.Println("Which one would you like to order?")
		bevId, err := strconv.Atoi(ScanInputString())
		if err != nil {
			panic(err)
		}
		//if drink not found or trying to buy alcohol while underage
		bev, ok := menu[bevId]
		if !ok || (bev.Alcohol && !isAdult) {
			fmt.Println("Beverage not found")
			continue
		}
		fmt.Println("How many would you like to buy?")
		quantity, err := strconv.Atoi(ScanInputString())
		if err != nil {
			panic(err)
		}
		total += bev.Price * float64(quantity)
		fmt.Printf("Current total = $%.2f\n", total)

		//generate query values and insert it to ordersDetailsQuery
		ordersDetailsQuery = fmt.Sprintf("%s (%v,%v,%v)", ordersDetailsQuery, orderId, bev.Id, quantity)

		//check if
		fmt.Println("Would you like to order more? (y/n)")
		char := ScanInputString()
		if char != "y" && char != "Y" {
			break
		}
		//add comma for multi value insertion
		ordersDetailsQuery = fmt.Sprintf("%s,", ordersDetailsQuery)
	}
	//print total
	fmt.Printf("Your total is $%.2f, Thank you for ordering!\n", total)
	//add semicolon
	ordersDetailsQuery = fmt.Sprintf("%s;", ordersDetailsQuery)
	//insert orderdetails
	cli.Handler.UserHandler.ExecContext(ctx, ordersDetailsQuery)
	cli.BackToMainMenu(customerId)
}
