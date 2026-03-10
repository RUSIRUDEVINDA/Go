package main

import "fmt"

func main() {
	i := 0
	fmt.Println("1st loop")
	for i < 5 {
		fmt.Printf("Value of i is %d\n", i)
		i++
	}

	fmt.Println("2nd loop")
	for i := 0; i < 5; i++ {
		fmt.Printf("Value of i is %d\n", i)
	}

	fmt.Println("3rd loop")
	for j := 0; j < 5; j++ {
		fmt.Printf("Value of j is %d\n", j)
	}

	names :=[]string{"Rusiru", "Devinda", "Ovindi", "Vimasha"}

	fmt.Println("4th loop")
	for k := 0; k < len(names); k++ {
		fmt.Println(names[k])
	}

	fmt.Println("5th loop")
	for index,value:=range names{
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	fmt.Println("6th loop")
	for _,value:=range names{
		fmt.Printf("the value is %v \n", value)
		value = "devid"
	}
}
