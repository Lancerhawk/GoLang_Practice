//Investment Calculator

package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Enter the value for the Inventment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter the value for the Years: ")
	fmt.Scan(&years)
	fmt.Print("Enter the value for the Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future Value of Investment:", futureValue)
	fmt.Println("Future Real Value of Investment (adjusted for inflation):", futureRealValue)
}
