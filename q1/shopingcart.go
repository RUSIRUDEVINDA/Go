package main

import "fmt"

func main() {
	var productName string = "Laptop"
	var productPrice float64 = 999.99
	var quantity int = 2
	var totalCost float64 = productPrice * float64(quantity)
	var isInStock bool = true

	fmt.Printf("Product Name: %s\n", productName)
	fmt.Printf("Price : %.2f\n", productPrice)
	fmt.Printf("Quantity : %d\n", quantity)
	fmt.Printf("Total Cost : %.2f\n", totalCost)
	fmt.Printf("In stock: %t\n", isInStock)

	// %s for string, how to fromate them? they are formatted using the %s verb in fmt.Printf. For example, to print a string variable called productName, you would use fmt.Printf("Product Name: %s\n", productName). The %s verb tells fmt.Printf to format the variable as a string.
	// %.2f for float, how to format them? they are formatted using the %.2f verb in fmt.Printf. For example, to print a float variable called productPrice with 2 decimal places, you would use fmt.Printf("Price: %.2f\n", productPrice). The %.2f verb tells fmt.Printf to format the variable as a float with 2 decimal places.
	// %d for int, how to format them? they are formatted using the %d verb in fmt.Printf. For example, to print an int variable called quantity, you would use fmt.Printf("Quantity: %d\n", quantity). The %d verb tells fmt.Printf to format the variable as an integer.
	// %t for bool, how to format them? they are formatted using the %t verb in fmt.Printf. For example, to print a bool variable called isInStock, you would use fmt.Printf("In stock: %t\n", isInStock). The %t verb tells fmt.Printf to format the variable as a boolean value (true or false).
}
