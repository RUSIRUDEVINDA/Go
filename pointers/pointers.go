package main

import "fmt"

func main() {
	var num int = 10
	fmt.Println("num : ", num)

	// Declaring a pointer to an integer
	var ptr *int
	// Assigning the address of num to the pointer
	ptr = &num

	fmt.Println("address : ", &num)
	fmt.Println("pointer : ", ptr)

	// dereferencing the pointer to get the value it points to
	fmt.Println(*ptr)
	num = 20
	fmt.Println(*ptr)

	fmt.Println("num : ", num)

	fmt.Println(&ptr)
}
