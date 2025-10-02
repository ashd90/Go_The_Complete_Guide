package main

import (
	"fmt"
	"os"
	// "strconv"
)

const accountBalanceFile = "balance.txt"

// func getBalanceFromFile() float64 {
// 	data, _ := os.ReadFile(accountBalanceFile)
// 	balanceText = string(data)
// 	balance, _ = strconv.ParseFloat(balanceText, 64)
// 	return balance

// }
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}
func main() {
	// var accountBalance = getBalanceFromFile()
	var accountBalance = 1000.0
	fmt.Println("Welcome to Bharat Bank")
	for i := 0; i < 3; i++ {

		fmt.Println("1. Check the Balance")
		fmt.Println("2. Deposit Cash")
		fmt.Println("3. Withdraw Cash")
		fmt.Println("4. Cancel Transaction")
		var choice int
		fmt.Print("Please Choose any one option: ")
		fmt.Scan(&choice)
		// fmt.Println("Your choice: ", choice)

		// checkBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your Account Balance is: ", accountBalance)
		case 2:
			fmt.Print("Enter Amount to Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Enter valid Amount to deposit")
				// return
				continue
			}
			var accountBalance = accountBalance + depositAmount
			fmt.Println("Your updated Balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Enter Amount to Withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Print("Enter valid amount to withdraw: ")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient Balance in your Account")
				fmt.Println("Your Balance is: ", accountBalance)
				continue
			}

			var accountBalance = accountBalance - withdrawAmount
			fmt.Println("Your Account Balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Your Transaction is cancelled! Remove Card")
			return
		}

		// if choice == 1 {
		// 	fmt.Println("Your Account Balance is: ", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Enter Amount to Deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)
		// 	if depositAmount <= 0 {
		// 		fmt.Println("Enter valid Amount to deposit")
		// 		// return
		// 		continue
		// 	}
		// 	var accountBalance = accountBalance + depositAmount
		// 	fmt.Println("Your updated Balance is: ", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("Enter Amount to Withdraw: ")
		// 	var withdrawAmount float64
		// 	fmt.Scan(&withdrawAmount)
		// 	if withdrawAmount <= 0 {
		// 		fmt.Print("Enter valid amount to withdraw: ")
		// 	}
		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("Insufficient Balance in your Account")
		// 		fmt.Println("Your Balance is: ",accountBalance)
		// 		return
		// 	}
		// 	var accountBalance = accountBalance -  withdrawAmount
		// 	fmt.Println("Your Account Balance is: ", accountBalance)
		// } else  {
		// 	fmt.Println("Your Transaction is cancelled! Remove Card")
		// 	break
		// }
	}

}
