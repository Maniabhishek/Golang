/*
useful to run operations after functions complete
the defer keyword can be used to execute code after a functions runs

	clean up resources, reset data etc
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 0, 30)
	fmt.Println(buffer)
	bytes, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		//file will close with deferred call
		return
	}
	fmt.Printf("%c\n", bytes)
	fmt.Println(bytes)
	//file will close with deffered bytes
}
