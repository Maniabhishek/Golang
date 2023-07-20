// Conversion characters
// Conversion characters tell Golang how to format different data types. Some of the most commonly used specifiers are:

// v – formats the value in a default format
// d – formats decimal integers
// g – formats the floating-point numbers
// b – formats base 2 numbers
// o – formats base 8 numbers
// t – formats true or false values
// s – formats string values

package main

import "fmt"

func main() {
	// str := "csdfd"
	fmt.Printf("here is %v value \n", map[string]interface{}{"a": 1})
	fmt.Printf("formats decimal integers %d \n", 1)
	//format floating point number
	fmt.Printf("format floating numbers %g \n", 1.2)
	fmt.Printf("format base 2 numbers %b \n", 4)
	fmt.Printf("format base 8 numbbers %o \n", 10)
	fmt.Printf("format true or false values %t \n", true)
	fmt.Printf("format string values %s\n", "string value")
}
