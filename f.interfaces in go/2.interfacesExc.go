package main

import "fmt"

type Lift int

const (
	smallLift = iota
	mediumLift
	largeLift
)

type LiftPicker interface {
	pickLift() Lift
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) String() string {
	return fmt.Sprintf("motorcycle: %v", string(m))
}

func (c Car) String() string {
	return fmt.Sprintf("motorcycle: %v", string(c))
}

func (t Truck) String() string {
	return fmt.Sprintf("motorcycle: %v", string(t))
}

func (m Motorcycle) pickLift() Lift {
	return smallLift
}

func (t Truck) pickLift() Lift {
	return largeLift
}

func (c Car) pickLift() Lift {
	return mediumLift
}

func sendToLift(p LiftPicker) {
	switch p.pickLift() {
	case smallLift:
		fmt.Printf("send %v to small lift\n", p)
	case mediumLift:
		fmt.Printf("send %v to medium lift\n", p)
	case largeLift:
		fmt.Printf("send %v to large lift\n", p)
	}
}

func main() {
	car := Car("rangeRover")
	truck := Truck("MountainCrusher")
	motorcycle := Motorcycle("yamaha")

	sendToLift(car)
	sendToLift(truck)
	sendToLift(motorcycle)
}
