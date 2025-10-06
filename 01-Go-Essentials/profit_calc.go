package main

import (
	"fmt"
)

func profit() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Please Enter Total Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Please Enter Total Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Please Enter Tax Rate: ")
	fmt.Scan(&taxRate)

	var ebt float64 = revenue - expenses
	var profit float64 = ebt * (1 - taxRate/100)
	var ratio = ebt / profit

	fmt.Printf("The Total Earnings Before Tax is: %.2f\n", ebt)
	fmt.Printf("The Total Earnings After Tax is: %.2f\n", profit)
	fmt.Printf("The ratio is: %.2f\n", ratio)

}
