package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Hello %v how you doing? \n", n)
}

func sayByee(n string) {
	fmt.Printf("Byee %v have a good day! \n", n)
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
		fmt.Println("-------------")
	}
}

func main() {
	// sayGreeting("Guptil")
	// sayByee("Kane")
	cycleNames([]string{"Desmod", "Jack"}, sayGreeting)
	cycleNames([]string{"Desmod", "Jack"}, sayByee)

	a1 := circleArea(5)
	a2 := circleArea(9)

	fmt.Println("Area of the a1 is: ", a1)
	fmt.Println("Area of the a2 is: ", a2)

	fmt.Printf("area of the a1 and a2 is %.3f and %.3f ", a1, a2)

}
