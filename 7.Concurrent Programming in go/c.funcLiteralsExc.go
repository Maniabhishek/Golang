package main

import (
	"fmt"
	"unicode"
)

type CountRune func(line string)

func calCulateTotalRune(lines []string, countRune CountRune) {
	for _, line := range lines {
		countRune(line)
	}
}

func main() {
	lines := []string{
		"1.golang is great,",
		"2.golang is great for type safety,",
		"3.golang is great for backend,",
		"4.let's learn golang.",
	}
	letter := 0
	punctuation := 0
	numbers := 0
	space := 0
	countRune := func(line string) {
		for _, char := range line {
			if unicode.IsLetter(char) {
				letter += 1
			} else if unicode.IsPunct(char) {
				punctuation += 1
			} else if unicode.IsSpace(char) {
				space += 1
			} else if unicode.IsNumber(char) {
				numbers += 1
			}
		}
	}

	calCulateTotalRune(lines, countRune)
	fmt.Printf("%v letters\n %v numbers \n %v punctuation, \n %v space", letter, numbers, punctuation, space)
}
