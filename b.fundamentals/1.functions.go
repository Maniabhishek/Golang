package main

import "fmt"

func greet(name string) string {
	fmt.Println("Hi ", name)
	return "Hi " + name
}

func fun() string {
	return "go is fun"
}

func add(num1 int, num2 int, num3 int) int {
	return num1 + num2 + num3
}

//num1 int , num2 int , num3 int this can be improved num1, num2, num3 int

func retrunAnyTwoNum() (int, int) {
	return 1, 2
}

func main() {
	fmt.Println(greet("rama"))
	fmt.Println(fun())
	fmt.Println(add(1, 2, 3))
}
