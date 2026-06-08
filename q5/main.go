// Write your answer here, and then test your code.

package main

import "fmt"
// Change these boolean values to control whether you see 
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

type CartItem struct{
    Name string
    Price float64
    Quantity int
}


// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []CartItem) float64 {
    var sum = 0.0
    for _, i := range cart { // this "item" means 
        sum += i.Price * float64(i.Quantity)
    }
    return sum
}

func main(){
	cart := []CartItem{}

	cart = append(cart, CartItem{Name: "apple", Price: 10, Quantity: 3}) 
	cart = append(cart, CartItem{Name: "orange", Price: 15, Quantity: 1})
	cart = append(cart, CartItem{Name: "grape", Price: 5, Quantity: 4})


	total := calculateTotal(cart)
	fmt.Println(total)
}


