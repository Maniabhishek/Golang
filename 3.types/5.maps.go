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

import "fmt"

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
