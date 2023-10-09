package main

import (
	"fmt"
	mypackage "mymodule/myPackage"
)

func main() {
	fmt.Println("this is main file")
	mypackage.MyPackage("function invoked from main file")
}
