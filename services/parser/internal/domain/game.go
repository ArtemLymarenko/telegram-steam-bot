package domain

import "math"

type Game struct {
	Id    string
	Name  string
	Price float64
	Url   string
}

func (g Game) PriceEquals(price float64) bool {
	return math.Abs(g.Price-price) < 0.01
}
