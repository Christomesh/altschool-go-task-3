package atm

import "fmt"

var balance float64 = 0.0

func CheckBalance() {
	fmt.Printf("Your balance $%.2f\n", balance)
}

func Withdraw() {
	var amount float64
	fmt.Println("How much do you want to withdraw: ")
	fmt.Scan(&amount)
	if amount > balance {
		fmt.Println("Insufficient balance")
		return
	}
	balance = balance - amount
	fmt.Printf("Success!: You have successfully withdraw $%.2f", amount)
	fmt.Println()

}

func Deposit() {
	var amount float64
	fmt.Println("How much do you want to deposit: ")
	fmt.Scan(&amount)
	balance = balance + amount
	fmt.Printf("Success!: You have successfully deposited $%.2f", amount)
	fmt.Println()
}
