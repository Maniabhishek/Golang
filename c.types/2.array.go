/*
creating an array
1.
var myArray [3]int
2. create and assign
myArray := [3]int{4,5,6}
3.
myArray := [...]int{3,4,6}
...will be replace by number of elements you assign
4.
myArray := [4]int{3,4,5}
at 4th indices value will be 0
*/

//elements not addressed in array initialization will be set to default values

//iteration
/*
good practice to assign the element to a variable during iteration

 Easier to read in large functions / nested loops
 myArray := [...]int{7,8,9}

 for i:=0 ; i< len(myArray) ; i++ {
	item := myArray[i]
	fmt.Println(item)
 }
*/

//array index out of bound
/*
myArray := [3]int
myArray[0] = 1
myArray[1] = 2
myArray[2] = 3
myArray[3] = 4 this will be compile time error as myArray has size of 3 only
*/
package main

import "fmt"

type Student struct {
	name      string
	graduated bool
}

type Product struct {
	name  string
	price float32
}

func checkIfGraduated(students [4]Student) {
	for i := 0; i < len(students); i++ {
		if students[i].graduated {
			fmt.Println(students[i].name, " is graduated")
		} else {
			fmt.Println(students[i].name, " is not graduated")
		}
	}
}

func main() {
	students := [...]Student{
		{name: "illy"},
		{name: "billy"},
		{name: "milly"},
		{name: "tilly"},
	}
	fmt.Println(students)
	students[0].graduated = true
	students[3].graduated = true
	checkIfGraduated(students)

	products := [...]Product{
		{name: "banana", price: 5},
		{name: "apple", price: 10},
		{name: "guava", price: 3},
		{name: "straberry", price: 8},
	}

	var sum float32

	for i := 0; i < len(products); i++ {
		sum += products[i].price
	}
	fmt.Println(sum)
}
