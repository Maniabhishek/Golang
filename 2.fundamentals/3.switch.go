/*
examples
x:=2
switch x {
case 1:
	fmt.Println(1)
case 2:
	fmt.Println(2)
default:
	fmt.Println("default")
}

*/

/*
conditional cases
switch x:= calculate(5);{
case x>10 :
	fmt.Println("greater than 10")
case x<10:
	fmt.Println("less than 10")
case x == 10:
	fmt.Println("eqaul to 10")
}
*/

//case list
/*
switch x {
case 1,2,3:
	....
case 4,5,6:
	....
}
*/

//fallthrough
/*
fallthrough will continue checking the next case
switch letter {
case ' ':
case 'a','e','i','o','u':
	fmt.Println("vowel")
	fallthrough
case 1,2,3,4:
	fmt.Println("numbers")
default:
	fmt.Println("default")
}

suppoer if space is provided case ' ' will be executed and exit and if case 'a','e'... is executed
then it will fallthrough and next case 1,2,3,4 will be executed irrespective of what value was provided
*/

package main

import "fmt"

func getAge() int {
	return 2
}

func main() {
	fmt.Print("main")
	switch age := getAge(); {
	case age == 0:
		fmt.Println("newborn")
	case age >= 1 && age <= 3:
		fmt.Print("toddler")
	case age <= 12:
		fmt.Print("child")
	case age <= 17:
		fmt.Print("teenager")
	default:
		fmt.Print("adult")
	}
}
