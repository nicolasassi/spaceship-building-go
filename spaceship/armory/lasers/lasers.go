package lasers

import (
	"fmt"
	"github.com/nicolasassi/spaceship-building-go/spaceship/armory"
	"time"
)

type Lasers struct {
	ammo int
}

func NewLasers(ammo int) *Lasers {
	return &Lasers{ammo: ammo}
}

func (l *Lasers) Shoot(x, y float32) error {
	if l.ammo >= 1 {
		fmt.Printf("Shooting with lasers at x:%v y:%v\n", x, y)
		l.ammo--
		return nil
	}
	return armory.RequiresReloading
}

func (l *Lasers) Reload(ammo int) {
	fmt.Println("Reloading lasers")
	for i := 0; i < ammo; i++ {
		time.Sleep(500 * time.Millisecond)
		l.ammo++
	}
	fmt.Println("lasers ready")
}
