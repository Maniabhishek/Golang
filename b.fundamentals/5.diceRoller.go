/*
print the sum of dice roll
print additional information in these cicrumstances:
	snake eyes when the total roll is 2 , and total dice is 2
	Lucky 7 when the total roll is 7
	Even : when the total roll is even
	odd when the total roll is odd
	the program must handle any number of dice , rolls and sides

	notes :
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand.Seed(3)
	// Seeding with the same value results in the same random sequence each run.
	// For different numbers, seed with a different value, such as
	// time.Now().UnixNano(), which yields a constantly-changing number.
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())

	dice, sides := 2, 12

	rolls := 1
	for i := 0; i < rolls; i++ {
		sum := 0
		for d := 0; d < dice; d++ {
			sum += rand.Intn(sides + 1)
		}
		fmt.Println("total rolled ", sum)
		switch sum := sum; {
		case sum == 2 && dice == 2:
			fmt.Println("snake eyes")
		case sum == 7:
			fmt.Println("lucky 7")
		case sum%2 == 0:
			fmt.Println("even")
		case sum%2 != 0:
			fmt.Println("odd")
		}
	}
}
