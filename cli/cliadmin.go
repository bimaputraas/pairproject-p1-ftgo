package cli

import (
	"fmt"
	"log"
	"strconv"
)

func (cli *Cli) MainMenuAdmin() {
	fmt.Printf("\nADMIN MENU\n\n")
	fmt.Println("[COMMAND]			-DESCRIPTION")
	fmt.Println("[view]				-view beverage")
	fmt.Println("[add]				-add beverage")
	fmt.Println("[delete]			-delete beverage")
	fmt.Println("[exit]				-exit")
	fmt.Printf("\nEnter your command : ")

	askAdminInput := ScanInputString()

	switch askAdminInput {
	case "view":
		cli.ViewBeveragesAdminInterface()
	case "add":
		cli.AddBeverageInterface()
	case "delete":
		cli.DeleteBeverageInterface()
	case "exit":
		fmt.Printf("\nAdmin access end, Thanks!\n\n")
		return
	default:
		fmt.Printf("Invalid Input\n\n")
		cli.MainMenuAdmin()
	}
}

func (cli *Cli) ViewBeveragesAdminInterface() {
	menu, err := cli.Handler.ViewBeverages()
	if err != nil {
		panic(err)
	}
	// print menu
	println("ID  Name                 Price       Alcohol")
	for _, bev := range menu {
		// print bev with string padding
		fmt.Printf("%-3s %-20s %-7s 	%s\n", fmt.Sprint(bev.Id), bev.Name, fmt.Sprintf("%.2f", bev.Price), fmt.Sprint(bev.Alcohol))
	}

	fmt.Println("")
	cli.MainMenuAdmin()
}

func (cli *Cli) AddBeverageInterface() {
	fmt.Printf("\n-ADD BEVERAGE-\n\n")
	// name
	fmt.Printf("\nPlease input beverage name: ")
	askName := ScanInputString()

	fmt.Printf("\nPlease input beverage price: ")
	askPriceStr := ScanInputString()
	askPrice, err := strconv.ParseFloat(askPriceStr, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nDoes beverage contains alcohol? (yes/no): ")
	askAlcohol := ScanInputString()

	var isContainsAlcohol bool
	switch askAlcohol {
	case "yes":
		isContainsAlcohol = true
	case "no":
		isContainsAlcohol = false
	default:
		isContainsAlcohol = false
	}

	cli.Handler.AddBeverage(askName, askPrice, isContainsAlcohol)
	fmt.Println("Beverage added to menu")
	cli.MainMenuAdmin()
}

func (cli *Cli) DeleteBeverageInterface() {
	fmt.Printf("\n-DELETE BEVERAGE-\n\n")

	// ask id
	fmt.Printf("\nPlease input beverage id: ")
	askIdStr := ScanInputString()
	askId, err := strconv.Atoi(askIdStr)
	if err != nil {
		log.Fatal(err)
	}

	err = cli.Handler.DeleteBeveragebyId(askId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Beverage deleted")
	cli.MainMenuAdmin()
}
