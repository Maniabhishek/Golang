/*
Receiver functions
modified function signature which allows dot notation
makes writing some types of functionality more convenient
allows simple muatation of existing structures
 similar to modifying a class variable in other language
*/
/*
regular function
type Coordinate struct {
	x int,
	y int,
}

func shiftBy(x,y int, coord *Coordinate{
	coord.x +=x
	coord.y +=y
})

coord := Coordinate{5,5}
shiftBy(1,1,&coord)
*/

/*
 Receiver function
 type Coordinate struct {
	x,y int
 }

 func (coord *Coordinate)shiftBy(x,y int){
	coord.x += x
	coord.y += y
 }

 coord := Coordinate{5,5}
 coord.shiftBy(1,2)
*/

/*
receiver functions provide the dot notations for structs

	create more convenient API's

Pointer receivers cannot modify a struct
value receivers cannot modify a struct
common to use pointer receivers
*/
package main

import "fmt"

type Space struct {
	taken bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].taken = true
}

func vacateSpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].taken = false
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].taken = true
}

func main() {
	parkingLot := ParkingLot{spaces: make([]Space, 10)}
	fmt.Println(parkingLot)
	occupySpace(&parkingLot, 1)
	fmt.Println(parkingLot)
	fmt.Println()
	parkingLot.occupySpace(3)
	fmt.Println()
	fmt.Println(parkingLot)

}
