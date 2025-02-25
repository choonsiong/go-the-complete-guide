package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const bankFileName = "bank_balance.txt"

func main() {
	//var balance float64
	isExit := false

	fmt.Println("-------------------------")
	fmt.Println("🏦 Welcome to Go Bank  🏦")
	fmt.Println("-------------------------")

	balance, err := readBalanceFromFile(bankFileName)
	if err != nil {
		fmt.Println("====> Read Balance Error:", err)
		panic(err)
	}

	for !isExit {
		switch menu() {
		case 1:
			checkBalance(balance)
		case 2:
			amount := getAmount()
			err := deposit(&balance, amount)
			if err != nil {
				fmt.Println("====> Deposit Error:", err)
			}
			checkBalance(balance)
		case 3:
			amount := getAmount()
			err := withdraw(&balance, amount)
			if err != nil {
				fmt.Println("====> Withdrawal Error:", err)
			}
			checkBalance(balance)
		default:
			err := writeBalanceToFile(balance, bankFileName)
			if err != nil {
				fmt.Println("====> Write Balance Error:", err)
			}
			fmt.Println("Thank you. See you again!")
			isExit = true
		}

		//choice := menu()
		//if choice == 1 {
		//	checkBalance(balance)
		//} else if choice == 2 {
		//	amount := getAmount()
		//	err := deposit(&balance, amount)
		//	if err != nil {
		//		fmt.Println("====> Deposit Error:", err)
		//	}
		//	checkBalance(balance)
		//} else if choice == 3 {
		//	amount := getAmount()
		//	err := withdraw(&balance, amount)
		//	if err != nil {
		//		fmt.Println("====> Withdrawal Error:", err)
		//	}
		//	checkBalance(balance)
		//} else if choice == 4 {
		//	fmt.Println("Thank you. See you again!")
		//	break
		//}
	}
}

func menu() int {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	var choice int
	fmt.Print("Enter your choice: ")
	_, _ = fmt.Scan(&choice)
	fmt.Println("Your choice is", choice)
	return choice
}

func checkBalance(balance float64) {
	fmt.Printf("====> Balance is: $%.2f\n\n", balance)
}

func getAmount() (amount float64) {
	fmt.Print("Enter your amount: ")
	_, _ = fmt.Scan(&amount)
	return
}

func deposit(balance *float64, amount float64) error {
	if amount < 0 {
		return errors.New("amount is negative")
	}
	*balance += amount
	return nil
}

func withdraw(balance *float64, amount float64) error {
	if *balance <= 0 {
		return errors.New("balance is zero")
	}
	*balance -= amount
	return nil
}

func readBalanceFromFile(filename string) (balance float64, err error) {
	balanceBytes, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// If balance file not found, create a new one with 0 balance
			return 0, os.WriteFile(filename, []byte(fmt.Sprint("0")), 0644)
		} else {
			return
		}
	}

	return strconv.ParseFloat(string(balanceBytes), 64)
}

func writeBalanceToFile(balance float64, filename string) error {
	return os.WriteFile(filename, []byte(fmt.Sprint(balance)), 0644)
}
