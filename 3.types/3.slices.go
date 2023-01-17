/*
slices are companion types that work with arrays
they enable a view into an array
views are dynamic an not fixed in size
functions can accept a slice as a function parameter
any size array can be operated upon via slice
*/

package main

import "fmt"

func main() {
	num := [4]int{1, 2, 3, 4}
	slice1 := num[1:3]
	fmt.Println(len(slice1))
}
