package main

import "fmt"

type Operation byte

const (
	Add Operation = iota
	Sub
	Mul
	Div
)

func (operation Operation) calculate(num1, num2 float32) float32 {
	switch operation {
	case Add:
		return (num1 + num2)

	case Sub:
		return (num1 - num2)
	case Mul:
		return (num1 * num2)
	case Div:
		return (num1 / num2)
	}
	panic("unhandled operation")
}

func main() {
	fmt.Println("adding 3 , 4 =", Add.calculate(3, 4))
	fmt.Println("subtracting 5 , 4 =", Sub.calculate(5, 4))
	fmt.Println("mul 3 , 4 =", Mul.calculate(3, 4))
	fmt.Println("divide 5 , 4 =", Div.calculate(5, 4))
}
