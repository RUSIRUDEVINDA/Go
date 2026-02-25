package main

import "fmt"

func main() {
	var itemPrice float64 = 45.25
	var itemQuantity int = 2
	var discountPercentage float64 = 10.0

	// cal the total cost
	var totalCost float64 = itemPrice * float64(itemQuantity)
	// cal the discount amount
	var totalDiscount float64 = (totalCost * discountPercentage / 100)
	// final cost
	var finalCost float64 = totalCost - totalDiscount

	fmt.Println("Total Cost : ", totalCost)
	fmt.Println("Total Discount : ", totalDiscount)
	fmt.Println("Final Cost : ", finalCost)
}
