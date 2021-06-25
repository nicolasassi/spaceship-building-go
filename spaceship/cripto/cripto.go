package cripto

const password = "687ds687ds6dsds879sdsdskh"

type Cripto struct {
	underBruteForceAttack bool
	tries                 int
}

func NewCripto() *Cripto {
	return &Cripto{
		underBruteForceAttack: false,
		tries:                 0,
	}
}

func (c *Cripto) reset() {
	c.underBruteForceAttack = false
	c.tries = 0
}

func (c *Cripto) GetUnderBruteForceAttack() bool {
	return c.underBruteForceAttack
}

func (c *Cripto) CheckPassword(passwdToCheck string) bool {
	if passwdToCheck != password {
		c.tries++

		if c.tries > 3 {
			c.underBruteForceAttack = true
		}

		return false
	}

	c.reset()

	return true
}
