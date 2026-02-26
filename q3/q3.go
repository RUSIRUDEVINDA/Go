package main

import "fmt"

func increment(value *int) {
	*value++
}

func zeroOut(value *int) {
	*value = 0
}

func main() {
	num := 8
	fmt.Println("num : ", num)

	increment(&num)
	fmt.Println("num after increment : ", num)

	zeroOut(&num)
	fmt.Println("num after zeroing out : ", num)
}

// M1 - increase number by 1
// M2 - convert number to 0
