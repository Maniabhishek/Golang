package main

import "fmt"

type Player struct {
	name              string
	health, maxHealth uint
	energy, maxEnergy uint
}

func (player *Player) addEnergy(amount uint) {
	player.energy += amount
	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	}
	fmt.Println()
	fmt.Println(player.name, ":", "energy :", player.energy, "health:", player.health)
}

func (player *Player) addEnergyDamage(amount uint) {
	if player.energy-amount > player.energy {
		player.energy = 0
	} else {
		player.energy -= amount
	}

	fmt.Println()
	fmt.Println(player.name, ":", "energy :", player.energy, "health:", player.health)
}

func (player *Player) addHealthDamage(amount uint) {
	if player.health-amount > player.health {
		player.health = 0
	} else {
		player.health -= amount
	}

	fmt.Println()
	fmt.Println(player.name, ":", "energy :", player.energy, "health:", player.health)
}

func (player *Player) addHealth(amount uint) {
	player.health += amount
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}
	fmt.Println()
	fmt.Println(player.name, ":", "energy :", player.energy, "health:", player.health)
}

func main() {
	player := Player{name: "robin", health: 0, maxHealth: 10, energy: 0, maxEnergy: 10}

	player.addHealth(5)
	player.addEnergy(5)

	player.addEnergyDamage(2)
	player.addHealthDamage(3)
}
