package main

import "fmt"

func main() {
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	var color = [3]string{"red", "green", "blue"}
	fmt.Println(color)

	color[0] = "yellow"
	fmt.Println(color)
}