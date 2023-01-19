package main

import "fmt"

func main() {
	slice := [...]string{"hello", "world", "!"}

	for i, element := range slice {
		fmt.Println(i, " ", element, ":")
		for _, char := range element {
			fmt.Printf(" %q\n", char)
		}
	}
}
