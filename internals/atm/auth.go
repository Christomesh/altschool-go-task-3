package atm

import (
	"errors"
	"fmt"
	"log"
)

var userPin string = "0000"

// To login the user
func Login() bool {
	fmt.Println("Your default pin is 0000")
	for {
		var pin string
		fmt.Println("Enter your Pin: ")
		fmt.Scan(&pin)
		if verifyPin(pin) {
			return true
		}
		break

	}
	return false
}

// To change pin
func ChangePin() {
	var oldPin string
	var newPin string
	// check if the old pin match
	fmt.Println("Enter old or default Pin: ")
	fmt.Scan(&oldPin)
	oldPin, err := validPin(oldPin)
	if err != nil {
		log.Println(err)
		return
	}
	if oldPin != userPin {
		log.Println("Old pin incorrect")
		return
	}
	// replace old pin with the pin
	fmt.Println("Enter new pin")
	fmt.Scan(&newPin)
	newPin, err = validPin(newPin)
	if err != nil {
		log.Println(err)
		return
	}
	userPin = newPin
	log.Println("Pin changed!")

}

// verify the user password
func verifyPin(pin string) bool {
	pin, err := validPin(pin)
	if err != nil {
		log.Println(err)
		return false
	}
	if pin != userPin {
		return false
	}
	return true
}

// Ensures the pin is four digits
func validPin(pin string) (string, error) {
	if len(pin) != 4 {
		return "", errors.New("\nInvalid pin: pin should be four digits")
	}
	return pin, nil
}
