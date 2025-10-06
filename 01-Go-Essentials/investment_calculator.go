package main

import (
	"fmt"
	"math"
)

// Golbal Variables
const inflationRate = 5.5

func main() {

	//const needs to be initialized whereas variable need not to be initialized

	// const inflationRate = 5.5
	var investmentAmount float64 //Value is not assigned
	// expectedReturns := 15.5 //Shorthand version where type is not explicitly mentioned
	var expectedReturns float64
	var years float64

	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount) //Takes input from the user whereas & represents the pointer

	fmt.Print("Enter expected Rate of Returns: ")
	fmt.Scan(&expectedReturns)

	fmt.Print("Investment Horizon: ")
	fmt.Scan(&years)

	futureValues, futureRealValue := calFutureValues(investmentAmount, expectedReturns, years)

	// var futureValue = investmentAmount * math.Pow(1 + (expectedReturns) / 100, years)
	// var futureRealValue = futureValue / math.Pow(1 + (inflationRate) / 100, years)

	// Refer the standard library for fmt here https://pkg.go.dev/fmt@go1.25.1

	fmt.Printf("The expected value of your total investment & returns: %.2f\n", futureValues) //Formatted print //.2f rounds off the decimal places to two
	fmt.Printf("The expected total investment & returns with adjusted inflation: %.2f\n", futureRealValue)

	//To add line breaks/multiline in the above Printf funtion use ` ` instead of ""
	//To save the above Printf output in a variable we need to use Sprintf(pass the arguments) funtion
	profit()
}

func calFutureValues(investmentAmount, expectedReturns, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+(expectedReturns)/100, years)
	rfv := fv / math.Pow(1+(inflationRate)/100, years)
	return fv, rfv
}
