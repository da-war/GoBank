package main

import (
	"fmt"

	"dawar.com/banikingApp/fileops"
)

const balanceFile = "balance.txt"

func main() {
	var balance, error = fileops.ReadBalanceFiles(balanceFile)
	if error != nil {
		fmt.Println("Error ")
		fmt.Println(error)
		// panic("Failed to read balance file")
	}

	var deposit float64
	var withdraw float64

	presentMenu()

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
		fileops.WriteBalanceFiles(balance, balanceFile)
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
			fileops.WriteBalanceFiles(balance, balanceFile)
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
