package engine

type Engine struct {
	Fuel int
	isOn bool
}

func NewEngine(fuel int, isOn bool) *Engine {
	if fuel == 0 {
		isOn = false
	}

	return &Engine{Fuel: fuel, isOn: isOn}
}

func (e *Engine) GetFuel() int {

	if e.Fuel == 0 {
		e.isOn = false
	}

	return e.Fuel
}

func (e *Engine) GetIsOn() bool {
	return e.isOn
}
