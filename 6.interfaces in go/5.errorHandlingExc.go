package main

import (
	"errors"
	"fmt"
)

type SomeObj struct {
	values []int
}

func (s *SomeObj) Get(index int) (int, error) {
	if index > len(s.values) {
		return 0, errors.New(fmt.Sprintf("no value at index %v", index))
	} else {
		return s.values[index], nil
	}
}

func main() {
	arr := SomeObj{}
	fmt.Println(arr.Get(1))
}
