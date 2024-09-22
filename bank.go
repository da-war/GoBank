package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func writeBalanceFiles(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func readBalanceFiles() (float64, error) {
	var balanceText, err = os.ReadFile(balanceFile)
	if err != nil {
		return 0, errors.New("Failed to read balance file")
	}

	balanceString := string(balanceText)
	balance, err := strconv.ParseFloat(balanceString, 64)
	if err != nil {
		return 0, errors.New("Failed to convert balance to float")
	}

	return balance, nil
}

func main() {
	var balance, error = readBalanceFiles()
	if error != nil {
		fmt.Println("Error ")
		fmt.Println(error)
	}

	var deposit float64
	var withdraw float64

	fmt.Println("Hello, World!")
	fmt.Println("Welcome to GoBank!")
	fmt.Println("What would you like to do?")
	fmt.Println("Please select an option:")
	fmt.Println("1. Create a new account")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Check balance")
	fmt.Println("5. Exit")

	var option int
	fmt.Scanln(&option)

	// if option == 1 {
	// 	fmt.Println("Creating Your Account")
	// 	fmt.Println("Your account has been created successfully")
	// } else if option == 2 {
	// 	fmt.Println("Depositing Money")
	// 	fmt.Print("Enter the amount you would like to deposit: ")
	// 	fmt.Scan(&deposit)
	// 	balance += deposit
	// 	fmt.Println("Deposit Successful")
	// 	fmt.Println("Your new balance is: ", balance)
	// } else if option == 3 {
	// 	fmt.Println("Withdrawing Money")
	// 	fmt.Print("Enter the amount you would like to withdraw: ")
	// 	fmt.Scan(&withdraw)
	// 	if withdraw > balance {
	// 		fmt.Println("Insufficient Funds")
	// 	} else {
	// 		balance -= withdraw
	// 		writeBalanceFiles(balance)
	// 		fmt.Println("Withdrawal Successful")
	// 		fmt.Println("Your new balance is: ", balance)
	// 	}

	// } else if option == 4 {
	// 	fmt.Println("Checking Balance...")
	// 	fmt.Println("Your balance is: ", balance)
	// } else if option == 5 {
	// 	fmt.Println("Exiting")
	// } else {

	// 	fmt.Println("Invalid Option")
	// }
	//changing above to switch

	switch option {
	case 1:
		fmt.Println("Creating Your Account")
		fmt.Println("Your account has been created successfully")
	case 2:
		fmt.Println("Depositing Money")
		fmt.Print("Enter the amount you would like to deposit: ")
		fmt.Scan(&deposit)
		balance += deposit
		writeBalanceFiles(balance)
		fmt.Println("Deposit Successful")
		fmt.Println("Your new balance is: ", balance)
	case 3:
		fmt.Println("Withdrawing Money")
		fmt.Print("Enter the amount you would like to withdraw: ")
		fmt.Scan(&withdraw)
		if withdraw > balance {
			fmt.Println("Insufficient Funds")
		} else {
			balance -= withdraw
			writeBalanceFiles(balance)
			fmt.Println("Withdrawal Successful")
			fmt.Println("Your new balance is: ", balance)
		}
	case 4:
		fmt.Println("Checking Balance...")
		fmt.Println("Your balance is: ", balance)
	case 5:
		fmt.Println("Goodbye!")
	default:
		fmt.Println("Invalid Option")

	}

}
