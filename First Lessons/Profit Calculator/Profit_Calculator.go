package main

import (
	"fmt"
)

func main() {
	var revenue, expenses float64
	taxRate := 0.3 // Default value can be given like this

	fmt.Print("Enter the value for Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the value for Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the value of the Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	fmt.Printf("Earnings Before Tax (EBT): %.2f\nProfit After Tax: %.2f\nProfit Ratio (EBT/Profit): %.2f", ebt, profit, ratio)
}
