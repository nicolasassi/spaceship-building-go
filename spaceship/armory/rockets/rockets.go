package rockets

import (
	"fmt"
	"github.com/nicolasassi/spaceship-building-go/spaceship/armory"
	"time"
)

type Rockets struct {
	ammo int
}

func NewRockets(ammo int) *Rockets {
	return &Rockets{ammo: ammo}
}

func (r *Rockets) Shoot(x, y float32) error {
	if r.ammo >= 1 {
		fmt.Printf("Shooting with rocket at x:%v y:%v\n", x, y)
		r.ammo--
		return nil
	}
	return armory.RequiresReloading
}

func (r *Rockets) Reload(ammo int) {
	fmt.Println("Reloading rocket")
	for i := 0; i < ammo; i++ {
		time.Sleep(2 * time.Second)
		r.ammo++
	}
	fmt.Println("rocket ready")
}
