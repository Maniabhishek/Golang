/*
slices are companion types that work with arrays
they enable a view into an array
views are dynamic an not fixed in size
functions can accept a slice as a function parameter
any size array can be operated upon via slice
*/

//slice syntax example
/*
	numbers := []int{1,2,3,4,5}
	slice1 := numbers[:]   o/p [1,2,3,4,5]
	slice2 := numbers[1:]  o/p [2,3,4,5]
	slice3 := numbers[:1]  o/p [1]
*/

//dynamic arrays
/*
	slices can be used to create arrays that can be exteneded
	the append() function can add additional elements
	numbers := []int{1,2,3}
	numbers = append(numbers,4,5,6)

	3 dots can be used to extend a slice with another slice
	part1 := []int{1,2,3}
	part2 := []int{4,5,6}
	combined := append(part1, part2...)
*/

//multidimensional slices
/*
board := [][]string{
	[]string{'a','b','c','d'},
	{'a','b','c','d'},
	{'a','b','c','d'},
	{'a','b','c','d'},
}

board[0][0] is a
*/

package main

import "fmt"

type Part string

func printAssembly(parts []Part) {
	for i := 0; i < len(parts); i++ {
		fmt.Println(parts[i])
	}
}

func main() {
	num := []int{1, 2, 3, 4}
	slice1 := num[1:3]
	fmt.Println(len(slice1))
	arr := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	fmt.Println(append(arr, arr2...))

	sliceMake := make([]int, 10)
	fmt.Println(sliceMake)

	assembly := []Part{"pipe", "belt", "wheels"}

	printAssembly(assembly)

	assembly = append(assembly, "rope", "rod")

	fmt.Println(assembly)
	assembly = assembly[3:]
	fmt.Println(assembly)
}
