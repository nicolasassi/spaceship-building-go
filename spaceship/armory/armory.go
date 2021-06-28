package armory

import "errors"

var RequiresReloading = errors.New("requires reloading")

type Armorier interface {
	Shoot(x, y float32) error
	Reload(ammo int)
}
