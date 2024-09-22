package main

import "fmt"

func main() {
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

	if option == 1 {
		fmt.Println("Creating a new account...")
	}
	if option == 2 {
		fmt.Println("Depositing money...")
	}
	if option == 3 {
		fmt.Println("Withdrawing money...")
	}
	if option == 4 {
		fmt.Println("Checking balance...")
	}
	if option == 5 {
		fmt.Println("Exiting...")
	}

}
