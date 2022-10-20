package main

import (
	"fmt"
	"github.com/Christomesh/altschool-go-task-3/internals/atm"
	"log"
	"os"
	"strings"
)

func main() {

	welecome()
	isAuth := atm.Login()
	if isAuth {
		menu()
	}
}

func welecome() {
	fmt.Println("********{ Welcome to the ATM CLI APP }************")
}
func performAnotherOperation() {
	var anotherOpt string
	fmt.Println("Do you want to perform another transaction (yes/no):  ")
	fmt.Scan(&anotherOpt)
	if strings.ToLower(anotherOpt) == "yes" {
		menu()
	}
	exitProgram()
}
func menu() {
	fmt.Println("Select Operation:")
	fmt.Println("1. Change Pin \t\t 2. Check Account Balance")
	fmt.Println("3. Withdrawal  \t 4. Deposit")
	fmt.Println("0. Exit program \t")

	var OpInput int
	_, err := fmt.Scan(&OpInput)
	if err != nil {
		log.Println("Error: Please select the correct menu item")
	}
	switch OpInput {
	case 1:
		atm.ChangePin()
		performAnotherOperation()
	case 2:
		atm.CheckBalance()
		performAnotherOperation()
	case 3:
		atm.Withdraw()
		performAnotherOperation()
	case 4:
		atm.Deposit()
		performAnotherOperation()
	case 0:
		exitProgram()
	default:
		fmt.Println("Invalid input")
		performAnotherOperation()
	}
}
func exitProgram() {
	fmt.Println("Thanks for banking with us.")
	fmt.Println("Goodbye👋")
	os.Exit(0)
}
