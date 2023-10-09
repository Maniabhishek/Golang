/*
Structures allow data to be stored in groups
 similar to a class in other programming languages
 each data point in the structure is called a field
 storing data in groups is usually more efficient
possible to associate functionality with structres
 Helps organize code and data
*/

// ddefining a strucutre

/*
type Sample struct {
	field string
	a, b int
}
*/

// instantiating a structure

/*
data := Sample{
	field: "world",
	a: 1,
	b: 2,
}
*/

/*
Default instantiation
any field not instanitiated during the instantiation will have default values

	data := Sample{}    this means that field will have "" , a and b will have 0
	data := Sample{a:5}
*/

/*
accessing fields
fields can be read from and written to
word := data.field
a,b := data.a , data.b

data.field = "hello"
data.a = 10
data.b = 50
*/

// anonymous structures
/*
	inline structs created using var will have default values
	shorthand version must have each other field defined

	this below snonymous way will have default value but in below case line 63 will not have default as we are create and assign thus we cant be messing with the values
	var sample struct {
		field string
		a, b int
	}

	sample.field = "hello"
	sample.a = 9

	another way
	sample := struct {
		field string
		a, b int
	} {
		"hello",
		1,2
	}
*/

package main

import "fmt"

type Employee struct {
	name   string
	age    int
	salary float32
	hired  bool
}

type Manager struct {
	Kam Employee
}

type Coordinate struct {
	x, y int
}

type Rectangle struct {
	a Coordinate
	b Coordinate
}

func main() {
	bob := Employee{"Bob", 25, 2300000, true}
	fmt.Println(bob)
	casey := Employee{
		name:   "casey",
		age:    27,
		salary: 2000000,
	}
	var (
		billa = Employee{name: "Bill", age: 25, salary: 2000003}
		elly  = Employee{name: "Elly", age: 25, salary: 20300003}
	)
	var pattrick Employee
	pattrick.age = 23
	pattrick.name = "Pattrick"
	pattrick.hired = true
	pattrick.salary = 5

	manager := Manager{bob}

	fmt.Println(casey)
	fmt.Println(elly)
	fmt.Println("billa", billa)
	fmt.Println(manager.Kam.hired)

	//calculate area of rectangle

	coordinate1 := Coordinate{1, 4}
	coordinate2 := Coordinate{2, 4}
	rectCoordinates := Rectangle{coordinate1, coordinate2}
	fmt.Println(calculateAreaOfRectangle(rectCoordinates))
}

func getWidth(rect Rectangle) int {
	return rect.b.x - rect.a.x
}

func getLength(rect Rectangle) int {
	return rect.b.y - 0
}

func calculateAreaOfRectangle(rect Rectangle) int {
	return getLength(rect) * getWidth(rect)
}
