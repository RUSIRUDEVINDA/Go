package main

import "fmt"

func main() {
	colors := []string{"red", "green", "blue"}

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for k := range colors {
		fmt.Println(colors[k])
	}

	value := 0
	sum := 0

	for value < 6 {
		sum += value
		fmt.Println(value)
		fmt.Println(sum)

		value++
	}
}
