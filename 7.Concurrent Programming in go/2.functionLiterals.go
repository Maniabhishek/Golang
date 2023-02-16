package main

import "fmt"

func add(lhs, rhs int) int {
	return lhs + rhs
}

func compute(lhs, rhs int, operation func(l, r int) int) int {
	fmt.Printf("numbers are %v and %v", lhs, rhs)
	return operation(lhs, rhs)
}

func main() {
	fmt.Printf("\nsum is %v", compute(3, 4, add))
}
