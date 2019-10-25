package car

type Speed float64
type Color string
type Wheels string

const (
	MPH Speed = 1.0
	KPH       = 1.6
)

const (
	BLUE   Color = "blue"
	RED          = "red"
	YELLOW       = "yellow"
)

const (
	SPORTS_WHEELS Wheels = "sports"
	STEEL_WHEELS         = "steel_wheels"
)

type Car struct {
	Speed  Speed
	Color  Color
	Wheels Wheels
}

type CarBuilder struct {
	Speed  Speed
	Color  Color
	Wheels Wheels
}

func (c *CarBuilder) SetSpeed(speed Speed) *CarBuilder {
	c.Speed = speed
	return c
}

func (c *CarBuilder) SetColor(color Color) *CarBuilder {
	c.Color = color
	return c
}

func (c *CarBuilder) SetWheels(wheels Wheels) *CarBuilder {
	c.Wheels = wheels
	return c
}

func (c *CarBuilder) Create() *Car {
	car := &Car{
		Speed:  c.Speed,
		Color:  c.Color,
		Wheels: c.Wheels,
	}
	return car
}
