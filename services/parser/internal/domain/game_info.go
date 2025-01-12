package domain

import (
	"math"
)

type GameInfo struct {
	GameId          int64
	Url             string
	ImageUrl        string
	InitialPrice    float64
	FinalPrice      float64
	DiscountPercent float64
}

func (g GameInfo) InitialPriceEquals(price float64) bool {
	return math.Abs(g.InitialPrice-price) < 0.01
}

func (g GameInfo) FinalPriceEquals(price float64) bool {
	return math.Abs(g.FinalPrice-price) < 0.01
}

func (g GameInfo) DiscountPercentEquals(discountPercent float64) bool {
	return math.Abs(g.DiscountPercent-discountPercent) < 0.01
}

func (gameInfo GameInfo) Validate() bool {
	if gameInfo.GameId < 0 {
		return false
	}

	if gameInfo.Url == "" {
		return false
	}

	if gameInfo.ImageUrl == "" {
		return false
	}

	if gameInfo.InitialPrice < 0 {
		return false
	}

	if gameInfo.FinalPrice < 0 {
		return false
	}

	if gameInfo.DiscountPercent < 0 {
		return false
	}

	return true
}
