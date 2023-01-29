/*
when you dont know how many arguments will be there inside a function then you can use variadics
*/

package main

import "fmt"

func addNumbers(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{7, 8, 9}
	newArr := append(a, b...)
	fmt.Println(addNumbers(newArr...))
}
