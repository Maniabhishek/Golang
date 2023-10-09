/*
the type of data expected by a function must be specified in the function parameters
dont always know the type ahead of time

	interfaces allow specifying behaviours of a type instead of the type itself
	this allows functions to operate on more than one type of data
*/

/*
interfaces are implicitly implemented
 when a type has all the receiver functions required by the interface , then it is considered implemented
functions operating on iterfaces should never accept a pointer to an interface
 caller etermines whether pointer or value is used
prefer multiple interfaces with a few functions over one large interface
*/

/*
when implementing a pointer receiver function , all function accepting the interface will only accept the pointers
 if self-modification is needed , implement all interface function as receiver functions for consistency

eg.,
type MyType int
func (m *MyType)Function1(){}
func (m MyType)Function1(x int)int{ return x}

then ...

func execute(i MyInterface){
	i.Function1()
}

m := MyType(1)
execute(m) this will not work will throw error
execute(&m)

*/

//so pointer receiver implementation
/*
type MyType int
func (m *MyType)Function1(){}
func (m *MyType)Function2(){}
above is not good way to implement

better make both as pointer receiver

*/

/*
let see an example
type Resetter interface {
	Reset()
}

type Coordinates struct {
	x,y int
}

type Player struct {
	health int
	position Coordinates
}

func (p *Player)Reset(r Resetter){
	p.health = 100
	p.position = Coordinates{0,0}
}

func Reset(r Resetter){
	r.Reset()
}

/*
access implementing type
it is sometimes needed to access the underlying type that implements an interface
 call functions , make modification , etc

func ResetWithPenalty(r Resetter){
	if player , ok := r.(Player); ok {
		player.health = 50
	}else{
		r.Reset()
	}
}
*/

package main

import "fmt"

type MyInterface interface {
	Function1()
	Function2(x int) int
}

type MyType int //implements MyInterface

func (m MyType) Function1() {
	fmt.Println("funciton1")
}
func (m MyType) Function2(x int) int {
	return x + x
}

func execute(i MyInterface) {
	i.Function1()
}

func mai2n() {
	//pass by value vs pointer
	x := MyType(1)
	// pass by pointer
	execute(&x)
	// pass by value
	execute(x)
}
