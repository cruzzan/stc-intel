package Entities

import "math"

type Club struct {
	Name    string `json:"name"`
	Id      int    `json:"id"`
	Classes []Class
}

func (c *Club) CountClasses() int {
	return len(c.Classes)
}

func (c *Club) CountFullyBookedClasses() int {
	fullyBooked := 0
	for _, class := range c.Classes {
		if !class.hasAvailableSpots() {
			fullyBooked++
		}
	}

	return fullyBooked
}

func (c *Club) GetFullyBookedPercentage() int {
	return int(math.Round((float64(c.CountFullyBookedClasses()) / float64(c.CountClasses())) * 100))
}
