package main

import "fmt"

const pi float32 = 3.14

func main() {
	// find the are of the circle
	var radius float32 = 5.0
	var area float32 = pi * radius * radius
	fmt.Println("area of the circle is: ", area)

	radius = 10.567
	area = pi * radius * radius
	fmt.Printf("area of the circle is: %.2f\n", area)
}
