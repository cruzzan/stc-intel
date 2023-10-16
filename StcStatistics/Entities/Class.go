package Entities

type Class struct {
	Name           string `json:"name"`
	AvailableSpots int    `json:"slots.leftToBookIncDropin"`
}

func (c *Class) hasAvailableSpots() bool {
	return c.AvailableSpots > 0
}
