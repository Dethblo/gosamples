package main

import (
	"example.com/gosamples/internal/calc"
	"fmt"
)

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calc.CalculateFutureValues(investmentAmount, expectedReturnRate, years)

	// basic approach
	//fmt.Println("Future Value: ", futureValue)
	//fmt.Println("Future Value (adjusted for inflation): ", futureRealValue)

	// formatted string approach
	//fmt.Printf("Future Value: %.2f\n", futureValue)
	//fmt.Printf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)

	// multi-line approach
	fmt.Printf(`
Future Value: %.2f
Future Value (adjusted for inflation): %.2f
	`, futureValue, futureRealValue)
}
