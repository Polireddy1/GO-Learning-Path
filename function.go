package main

import (
	"fmt"
)


/* func amne(parameters) (returnTypes) {
   function body
}                 */


func calculateBill(price, no int) int {
	var totalPrice = price * no
	//totalPrice:=price*no
	return totalPrice                                  //FUNCTIONS 
}

func main() {
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)
}
