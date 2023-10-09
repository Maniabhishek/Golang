/*
const is like a variable , but unchanging
common to make groups of constants
iota keyword can be use to automatically assign values

*/

package main

import "fmt"

const (
	zero  = 0
	one   = 1
	two   = 2
	three = 3
	four  = 4
)

//or we can use iota for initialisation

const (
	zero1  = iota
	one1   // 1
	two2   // 2
	three3 // 3
	four4  // 4
)

//long form
const (
	m0 = iota
	m1 = iota
	m2 = iota
	m3 = iota
	m4 = iota
)

//short form
const (
	n0 = iota
	n1
	n2
	n3
	n4
)

//skip a value

const (
	s0 = iota
	_
	_
	s3
	s4
	s5
)

//start with a specific value
const (
	t0 = iota + 3
	t1
	t2
	t3
)

//iota enumeration pattern

type Direction byte

const (
	North Direction = iota
	East
	West
	South
)

func (d Direction) Stirng() {
	switch d {
	case North:
		fmt.Println("north")
	case South:
		fmt.Println("south")
	case East:
		fmt.Println("east")
	case West:
		fmt.Println("west")
	}
}

//shorter way for above with using the switch

func (d Direction) printDirection() string {
	return []string{"North", "East", "South", "West"}[d]
}

func main() {
	fmt.Println(n1, n4, s5, t3)
	n := North

	n.Stirng()

	s := South
	s.Stirng()
	fmt.Println(n.printDirection())
}
