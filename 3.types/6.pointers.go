/*
function calls in go are pass by value
a copy of each function argument is made , regardless of size
potentially slow for large data structures
more difficult to manage program state
this can be changes by using pointers
*/

/*
the value of the variable itself is memory address
accessing the data requires dereferencing the pointer
this allows changing values that exist elsewhere in the program
*/

/*
creating pointers
value := 10
var valuePtr *int
valuePtr = &value  (&value means we created pointer from variable )

shorter way
value := 10
valuePtr := &value

*/

/*
asterisk (*) when use with the pointer will dereference the pointer
this provides access to the actual data it points to
func increment(x *int){
	*x += 1
}

value := 1
increment(&value)

*/

package main

import "fmt"

type Counter struct {
	count int
	num   int
}

func increment(counter *Counter) {
	counter.count += 1
	fmt.Println("counter is ", *counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	fmt.Println("activating...")
	*tag = Active
}

func deActivate(tag *SecurityTag) {
	fmt.Println("deactivating...")
	*tag = Inactive
}

func checkout(items []Item) {
	fmt.Println("checking out...")
	for i := 0; i < len(items); i++ {
		deActivate(&items[i].tag)
	}
}

func printItems(items []Item) {
	fmt.Println("printing items")
	for i := 0; i < len(items); i++ {
		fmt.Println(i, items[i].name, items[i].tag)
	}
}

func main() {
	counter := Counter{}
	hello := "hello"
	world := "world"
	replace(&hello, "hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[0], "go", &counter)
	fmt.Println(phrase)
	
	//excercises
	
	fmt.Println("println")
	items := []Item{
		{name: "kLaptop", tag: Active},
		{name: "kMobile", tag: Active},
		{name: "PersonalLaptop", tag: Active},
		{name: "PersonalMobile", tag: Active},
	}

	deActivate(&items[0].tag)
	deActivate(&items[2].tag)

	printItems(items)

	checkout(items)

	printItems(items)
}
