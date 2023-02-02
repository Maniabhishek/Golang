/*
go has no exception
errors are returned as the last return value from from a function
  encodes failure as part of the function signature
    simple to determine if a function can fail
  return nil if no error occurred
errors implement the error interface from std
  one function to implement : Error() string
*/

/*
	example
	import "errors"

	func handleError(num1 int, num2 int)(int, error){
		if(num1>num2){
			return 0, error.New("num2 should be greater than num1")
		}else {
			return num2-num1,nil
		}
	}
*/

//error interface
/*

type error interface{
	Error() string
}
*/

//implementation   always implement error as receiver function
//helps prevent comparison problems if error is inspected
/*
type DivError struct{
	a, b int
}

func (d *DivError)Error()string{
	reutrn fmt.Sprintf("cannot divide by zero %d / %d",d.a,d.b)
}

*/

//use errors.Is() to determine if an error contains a specific type

package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

type DivError struct {
	a, b int
}

type UserError struct {
	msg string
}

func (d *DivError) Error() string {
	return fmt.Sprintf("cannot divide by zero %d / %d", d.a, d.b)
}

func (u *UserError) Error() string {
	return fmt.Sprintf("user error %v", u.msg)
}

func divide(a, b int) (float32, error) {
	if b == 0 {
		return 0, &DivError{a, b}
	} else {
		return float32(a / b), nil
	}
}

func someFunc(str string) (string, error) {
	if str == "" {
		return "", &UserError{"empty error"}
	} else {
		return str, nil
	}
}

func main() {
	val, err := divide(2, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

	v, err := someFunc("")
	if err != nil {
		if errors.Is(err, &UserError{"empty error"}) {
			fmt.Println("input error", err)
		} else {
			fmt.Println("other error", err)
		}
	}

	var thisError *UserError

	if err != nil {
		if errors.As(err, &thisError) {
			fmt.Println("user error", err)
		} else {
			fmt.Println("other error", err)
		}
	}
	fmt.Println(v)
}
