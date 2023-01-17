/*
for i:=0; i<10 ; i++ {
	....
}

for while
for i<10{
	...
	i++
}

for infinite
for {
	.....
	if somethingBreaks
		break
}

continue to skip iteration
*/

package main

import "fmt"

func main() {
	fmt.Println("main ")
	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
	}
	for sum > 10 {
		sum -= 5
		fmt.Println("sum is ", sum)
	}

	for i := 1; i <= 50; i++ {
		if i%3 == 0 {
			fmt.Println(i, ". fizz")
		} else if i%5 == 0 {
			fmt.Println(i, ". buzz")
		} else if i%5 == 0 && i%3 == 0 {
			fmt.Println(i, ". fizzbuzz")
		}
	}

	// or above can be improved
	for i := 1; i <= 50; i++ {
		divisibleBy3 := i%3 == 0
		divisibleBy5 := i%5 == 0

		if divisibleBy3 && divisibleBy5 {
			fmt.Println(i, "fizbuzz")
		} else if divisibleBy3 {
			fmt.Println(i, "fiz")
		} else if divisibleBy5 {
			fmt.Println(i, "buzz")
		} else {
			fmt.Println(i)
		}
	}
}
