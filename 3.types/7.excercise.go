//program to manage a lending of bikes

//requirements
/*
the library must have bike and memebers , and must include:
- which bike have been checked out
- what time bike were checked out
- what time bike were returned
Perform the following
- Add at least 4 books and at least 3 members to the library
- check out a book
- check in a book
- print out initial library information , and after each change
there must only ever be one copy of the library in memory at any time
*/

package main

import (
	"fmt"
	"time"
)

type BikeName string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	bikes map[BikeName]LendAudit
}

type BikeEntry struct {
	total  int
	lended int
}

type StoreHouse struct {
	member map[Name]Member
	bikes  map[BikeName]BikeEntry
}

func printMemberAudit(member *Member) {
	for bikeName, audit := range member.bikes {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", bikeName, audit.checkOut.String(), "to", returnTime)
	}
}

func printMemberAudits(storeHouse *StoreHouse) {
	for _, member := range storeHouse.member {
		printMemberAudit(&member)
	}
}

func printStoreHouseBikes(storeHouse *StoreHouse) {
	for bikeName, bike := range storeHouse.bikes {
		fmt.Println(bikeName, "total", bike.total, "lended", ":", bike.lended)
	}
}

func checkOutBook(storeHouse *StoreHouse, bikeName BikeName, member *Member) bool {
	bike, found := storeHouse.bikes[bikeName]

	if !found {
		fmt.Println("bike not available")
		return false
	}
	if bike.total == bike.lended {
		fmt.Println("all the bike have been lended")
		return false
	}

	bike.lended += 1
	member.bikes[bikeName] = LendAudit{
		checkOut: time.Now(),
	}
	storeHouse.bikes[bikeName] = bike
	return true
}

func returnBike(storeHouse *StoreHouse, bikeName BikeName, member *Member) bool {
	bike, found := storeHouse.bikes[bikeName]

	if !found {
		fmt.Println("bike not part of storehuse")
	}

	audit, found := member.bikes[bikeName]
	if !found {
		fmt.Println("member did not lend this bike")
		return false
	}

	bike.lended -= 1
	storeHouse.bikes[bikeName] = bike

	audit.checkIn = time.Now()

	member.bikes[bikeName] = audit
	return true
}

func main() {
	storeHouse := StoreHouse{
		bikes:  make(map[BikeName]BikeEntry),
		member: make(map[Name]Member),
	}

	storeHouse.bikes["yamaha"] = BikeEntry{total: 5}
	storeHouse.bikes["enfield"] = BikeEntry{total: 3}
	storeHouse.bikes["apache"] = BikeEntry{total: 2}
	storeHouse.bikes["honda"] = BikeEntry{total: 4}
	storeHouse.bikes["bmw"] = BikeEntry{total: 1}

	storeHouse.member["tin"] = Member{name: "tin", bikes: make(map[BikeName]LendAudit)}
	storeHouse.member["min"] = Member{name: "min", bikes: make(map[BikeName]LendAudit)}
	storeHouse.member["cin"] = Member{name: "cin", bikes: make(map[BikeName]LendAudit)}
	storeHouse.member["kin"] = Member{name: "kin", bikes: make(map[BikeName]LendAudit)}

	fmt.Println("\ninitial:")
	printStoreHouseBikes(&storeHouse)

	member := storeHouse.member["tin"]

	fmt.Println("after checkout bmw")
	checkOutBook(&storeHouse, "bmw", &member)
	fmt.Println()
	printStoreHouseBikes(&storeHouse)

	fmt.Println("checked out ...")
	fmt.Println()
	member2 := storeHouse.member["min"]

	checkOutBook(&storeHouse, "honda", &member2)
	printStoreHouseBikes(&storeHouse)

	fmt.Println()

	fmt.Println("audits...")

	printMemberAudits(&storeHouse)
}
