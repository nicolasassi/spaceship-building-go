package main

import (
	"github.com/nicolasassi/spaceship-building-go/spaceship/armory/rockets"
	"math/rand"
)

func Attack(r *rockets.Rockets) {
	var x, y float32
	for i := 0; i < 3; i++ {
		x, y = generateRandomPoints()
		if err := r.Shoot(x, y); err != nil {
			if err == rockets.RequiresReloading {
				r.Reload(1)
			}
		}
	}
}

func main() {
	r := rockets.NewRockets(1)
	Attack(r)
}

func generateRandomPoints() (float32, float32) {
	return float32(rand.Intn(30-0) + 0), float32(rand.Intn(30-0) + 0)
}
