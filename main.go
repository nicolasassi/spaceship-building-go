package main

import (
	"github.com/nicolasassi/spaceship-building-go/spaceship/armory"
	"github.com/nicolasassi/spaceship-building-go/spaceship/armory/lasers"
	"github.com/nicolasassi/spaceship-building-go/spaceship/armory/rockets"
	"math/rand"
)

func attack(arms ...armory.Armorier) {
	var x, y float32
	for i := 0; i < 3; i++ {
		x, y = generateRandomPoints()
		for _, arm := range arms {
			if err := arm.Shoot(x, y); err != nil {
				if err == armory.RequiresReloading {
					arm.Reload(1)
				}
			}
		}
	}
}

func main() {
	r := rockets.NewRockets(1)
	l := lasers.NewLasers(1)
	attack(r, l)
}

func generateRandomPoints() (float32, float32) {
	return float32(rand.Intn(30-0) + 0), float32(rand.Intn(30-0) + 0)
}
