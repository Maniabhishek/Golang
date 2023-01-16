// all data in programs consists of binary numbers (0 or 1)
// A data type is a way that the program can interpret the binary numbers
//Numbers letters  and words are all different type eg 2, A , Hello
//go is statically typed language
// data type must be provided by the programmer
//go uses type inference to determine what type of data is working with
// data types only need to be provided in specific circumstances
// can always specify the type if desired
//compiler error if wrong type is used
//all primitve data types in go are numeric
//type indicated in code is a convention
//it's possible that the data is invalid for the given type
// only applies when working with user input or manually manipulating the binary data

/*
signed integer types
int8  -128 to 127   (int 8 2^8 256 which mean if signed(-ve and +ve) then 256/2 i.e., -128 to 137 )
int16  -32768 to -32767
int 32
int 64
*/

/*
unsigned integer types
uint8 min 0 max 255
byte  min 0 max 255
uint16 min 0 65535
uintptr 0 <pointer size> (unsigned integer pointer)
*/

/*
other data types
Data Type    Description
float32      32-bit floating point
float64      64-bit floating point
complex64    32-bit floating point real and imaginary
complex128   64 bit floating point real and imaginary
bool         true and false
*/

/*
TYPE ALIASES
possible to create type aliases
same in every way to another type , just a different name
useful for providing indication of what kind of data is being utilized
type userId int
type Direction byte
type Speed float64
type Velocity Speed
*/

/*
Type conversions
converting between types can be done with parantheses
type UserId int
type Speed float64
UserId(5)
Speed(88.3)
*/

/*
string and runes
Runes
Text is represented using the rune type
similar to char in many other programming languages
Rune is an alias for int32(32-bit integer)
Always a number : will print numeric value unless proper formatting is specified
A rune can represent any symbol
letters , numbers, emoji, etc
String
A string is the data type for storing multiple runes
string are just an array of bytes and a string length
there is no null termination with go String
 bytes are not symbols
 special iteration required to retrieve runes/symbols

Creation
Runes: 'a', 'R', '7', '\n',
string "Hello there"
Raw literal : "let code in go"
*/

/*
Everyday go commands
build : builds the project & emits an executable binary
build-race: checks for concurrency problems
run: runs the projects directly; no output executable
mod: manages modules & dependecies
mod tidy : updates dependencies
test : run the projects test suite
fmt : formats all source files (usually automated with IDE)
*/

/*
creating variables
var keyword in Golang is used to create the variables of a particular type having a proper name and initial value. Initialization is optional at the time of declaration of variables using var keyword that we will discuss later in this article.
Syntax:

var identifier type = expression
var num int = 200
As you know that Go is a statically typed language but it still provides a facility to remove the declaration of data type while declaring a variable as shown in below syntax. This is generally termed as the Type Inference.

Syntax:

var identifier = initialValue

Example:

var geek1 = 200
Multiple variable declarations using var Keyword
Declaring multiple variables using var keyword along with the type:
var geek1, geek2, geek3, geek4 int

Declaring multiple variables using var keyword along with the type and initial values:
var geek1, geek2, geek3, geek4 int = 10, 20, 30, 40

You can also use type inference(discussed above) that will let the compiler to know about the type i.e. there is an option to remove the type while declaring multiple variables.
Example:

var geek1, geek2, geek3, geek4 = 10, 20, 30.30, true

You can also use multiple lines to declare and initialize the values of different types using a var keyword as follows:
Example:

var(
     geek1 = 100
     geek2 = 200.57
     geek3 bool
     geek4 string = "GeeksforGeeks"
)

While using type during declaration you are only allowed to declare multiple variables of the same type. But removing type during declarations you are allowed to declare multiple variables of different types.
*/

package main

import "fmt"

func main() {
	var num1 int = 23
	fmt.Println(num1)
	var num2, num3 int = 1, 2
	fmt.Println(num2, num3)
	var num4, string, num5 = 1, "abc", 4
	fmt.Println(num4, string, num5)
	main2()
	createAndAssignVar()
}

/*
Important Points about var keyword:

During the declaration of the variable using var keyword, you can either remove type or = expression but not both. If you will do, then the compiler will give an error.
If you removed the expression then the variable will contain the zero-value for numbers and false for booleans “” for strings and nil for interface and reference type by default. So, there is no such concept of an uninitialized variable in Go language.
*/

func main2() {

	// Variable declared but
	// no initialization
	var geek1 int
	var geek2 string
	var geek3 float64
	var geek4 bool

	// Display the zero-value of the variables
	fmt.Printf("The value of geek1 is : %d\n", geek1)

	fmt.Printf("The value of geek2 is : %s\n", geek2)

	fmt.Printf("The value of geek3 is : %f\n", geek3)

	fmt.Println("The value of geek4 is : %t", geek4)

}

/*
create and assign
num1 := 3
a,b := 1, "sample"
*/

func createAndAssignVar() {
	x, err, newErr := 1, "this is error", "this is new error"

	y, err, newErr := 2, "this is another error", "this is another new error"

	fmt.Println(x, y, err, newErr)

}
