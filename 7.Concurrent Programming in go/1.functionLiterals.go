/*
function literals provide a way to define a function within a function
possible to assign a function literals to variables
they can be passed to a function as parameters
 more dynamic code
also known as closures or anonymous functions
 closures allow data to be encapsulated within
*/

// anonymous function
package main

import "fmt"

func helloWorld() {
	fmt.Printf("hello,")
	world := func() {
		fmt.Printf("world! \n")
	}
	world()
	world()
	world()
	world()
}

func customMsg() func(msg string) {
	return func(msg string) {
		fmt.Printf("%.*s\n", len(msg), "---")
		fmt.Println(msg)
		fmt.Printf("%.*s\n", len(msg), "---")
	}
}

func main() {
	customMsg()("hello")
	discount := 0.1
	discountFn := func(subtotal float64) float64 {
		if subtotal > 100 {
			discount += 0.1
		}

		if discount > 0.3 {
			discount = 0.3
		}
		return discount
	}
	totalPrice := calculatePrice(200, discountFn)
	fmt.Println(totalPrice)
}

func calculatePrice(
	subTotal float64,
	discountFn func(subTotal float64) float64) float64 {
	return (subTotal - (subTotal * discountFn(subTotal)))
}

// as we see in the function calculate price line 53 it has become pretty verbose so we can use type alias as shown below

type DiscountFunc func(subTotal float64) float64
