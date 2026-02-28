package main

import "fmt"

func main() {
	// create a 2D array of integers
	var numbers = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(numbers)

	// take 8
	fmt.Println(numbers[2][1])
}