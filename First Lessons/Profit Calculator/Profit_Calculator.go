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
	fmt.Println("Earnings Before Tax (EBT):", ebt)
	fmt.Println("Profit After Tax:", profit)
	fmt.Println("Profit Ratio (EBT/Profit):", ratio)
}
