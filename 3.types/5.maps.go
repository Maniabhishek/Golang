/*
stores daa in key pairs
extremelyhigh performance when the key is known
unordered data is stored in random order
*/

//create map:
/*
myMap := make(map[string]int)
myMap := map[string]int{
	"item 1": 1,
	"item 2": 2,
	"item 3": 3,
}
*/

// insert read delete check existence
/*
	myMap = make(map[string]int)

	insert
	myMap["first number"] = 1

	read
	firstNum := myMap["first number"]
	missing := myMap["second number"] // this will have default value

	delete
	delete(myMap, "first number")

	check existence
	num, found := myMap["second number"]
	if !found {
		....
	}

	//iteration
	for key, value := range myMap {
		//..
	}
*/

package main

import (
	"fmt"
	"strconv"
)

const (
	online      = 0
	offline     = 1
	maintenance = 2
	retired     = 3
)

func main() {
	myMap := make(map[string]int)
	myMap["item 1"] = 1
	fmt.Println(myMap)

	missing := myMap["price"]
	fmt.Println(missing)

	delete(myMap, "item 1")

	item1, found := myMap["item 1"]
	if !found {
		fmt.Println(item1, "not found", found)
	}
	printServer()
	structWithinMap()
	fmt.Println(checkMapIfKeyExistInMap())
	fmt.Println(checkIfStockExists())
	fmt.Println(checkIfStockExists2())
}

func printServerStatus(server map[string]int) {
	fmt.Println("there are", len(server), "servers")
	serverStatus := make(map[int]int)
	for _, status := range server {
		switch status {
		case online:
			serverStatus[online] += 1
		case offline:
			serverStatus[offline] += 1
		case maintenance:
			serverStatus[maintenance] += 1
		case retired:
			serverStatus[retired] += 1
		}
	}
	fmt.Println(serverStatus)
	fmt.Println(serverStatus[online], "server are online")
	fmt.Println(serverStatus[offline], "server are offline")
	fmt.Println(serverStatus[retired], "server are retired")
	fmt.Println(serverStatus[maintenance], "server are maintenance")

}

func printServer() {
	servers := []string{"darkstar", "aiur", "omicron", "w359"}
	serverStatus := make(map[string]int)

	for _, server := range servers {
		serverStatus[server] = online
	}
	printServerStatus(serverStatus)
}

// There's one little gotcha to look out for here: we can't assign to a struct field within a map. For example, suppose we had a type menuItem like this:
type menuItem struct {
	price float64
}

func structWithinMap() {
	var menu = map[string]menuItem{
		"beans": menuItem{
			price: 0.49,
		},
	}

	// menu["beans"].price = 0.25 // not allowed
	beans := menu["beans"]
	beans.price = 0.25
	menu["beans"] = beans
	fmt.Println(menu["beans"])
}

// Checking whether a key exists in the map
// Sometimes it's useful to be able to tell whether or not a Go map key exists, and there's a special syntax for that:
func checkMapIfKeyExistInMap() string {

	menu := map[string]menuItem{
		"mobile": menuItem{
			price: 1234,
		},
	}

	fmt.Println(menu)

	// syntax for checking if exists
	res, ok := menu["beans"]

	if ok {
		fmt.Println("beans exists ")
		return strconv.Itoa(int(res.price))
	}
	return "key does not exist"
}

// Boolean maps
// Suppose we have a map called inStock which represents our current stock:

func checkIfStockExists() string {
	var stock = map[string]bool{
		"pizza":   true,
		"burger":  true,
		"biryani": false,
	}

	if _, ok := stock["chicken"]; ok {
		return fmt.Sprintf("Yes, %s is in stock", "chicken")
	}
	return "chicken not available"
}

//	This is a little lame, though. Since we know that maps return the zero value for mising keys,
// 	and we also happen to know that the zero value of bool is false, we can write simply:

func checkIfStockExists2() string {
	var stock = map[string]bool{
		"pizza":   true,
		"burger":  true,
		"biryani": false,
	}

	if stock["pizza"] {
		return fmt.Sprintf("Yes, %s is in stock", "pizza")
	}
	return "pizza not available"
}
