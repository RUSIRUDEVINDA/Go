package main

import "fmt" // fmt use for formatted I/O

func main() {
	fmt.Print("hello world")
	fmt.Print("\thello world")
	fmt.Print("\thello world")

	var userName string = "Rusiru Devinda"
	fmt.Println("\t" + userName)

	var userScore float32 = 25.6
	var userBalance float64 = -26.89676

	fmt.Println(userScore, userBalance)
	fmt.Printf("Type: %T", userName)

	var isUserFound bool = true
	fmt.Println(isUserFound)

	// multiple variable declaration
	var firstName, secondName = "Rusiru", "Devinda"
	fmt.Println(firstName, secondName)

	// type inference, Go infers the type from the value on the right
	newUserName := "Rusiru New"
	fmt.Println(newUserName)
	fmt.Printf("%T", newUserName)
}
