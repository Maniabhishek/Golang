/*
statement initialisation
if i:=5; i<10{
	//do something
}

if rank := getUserRank(); rank == "admin" {

}else if rank == "manager" {

} else {

}
*/

package main

import "fmt"

func average(a, b, c int) float32 {

	total := a + b + c
	return float32(total / 3)
}

func accessGranted(day int, role string) string {
	if role == "Admin" || role == "manager" {
		return "Permission allowed"
	} else if (day == 5 || day == 6) && role == "contractor" {
		return "Permission allowed"
	} else {
		return "Permission denied"
	}
}

func main() {
	var quiz1, quiz2, quiz3 = 1, 2, 4
	if quiz1 > quiz2 {
		fmt.Println("quiz1 has higher value")
	} else if quiz2 > quiz1 {
		fmt.Println("quiz2 has higher value")
	}

	if quiz3 > quiz1 && quiz3 > quiz2 {
		fmt.Println("quiz3 is highest")
	}

	if average(1, 2, 3) > 1 {
		fmt.Println("good")
	} else {
		fmt.Println("bad")
	}

}
