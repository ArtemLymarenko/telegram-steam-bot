package game

import (
	"errors"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain"
	"math"
)

type Url string

func (url Url) Validate() domain.ValidationError {
	if url == "" {
		return errors.New("url must not be empty")
	}
	return nil
}

type ImageUrl string

func (imageUrl ImageUrl) Validate() domain.ValidationError {
	if imageUrl == "" {
		return errors.New("imageUrl must not be empty")
	}
	return nil
}

type Price float64

func (price Price) Validate() domain.ValidationError {
	if price < 0 {
		return errors.New("initialPrice must be greater than 0")
	}
	return nil
}

func (price Price) EqualsTo(to Price) bool {
	return math.Abs(float64(price-to)) < 0.01
}

type DiscountPercent float64

func (discountPercent DiscountPercent) Validate() domain.ValidationError {
	if discountPercent < 0 {
		return errors.New("discountPercent must be greater than 0")
	}
	return nil
}

func (discountPercent DiscountPercent) EqualsTo(to DiscountPercent) bool {
	return math.Abs(float64(discountPercent-to)) < 0.01
}

type Info struct {
	GameId          Id
	Url             Url
	ImageUrl        ImageUrl
	InitialPrice    Price
	FinalPrice      Price
	DiscountPercent DiscountPercent
}

func (gameInfo Info) Validate() domain.ValidationError {
	if err := gameInfo.GameId.Validate(); err != nil {
		return err
	}

	if err := gameInfo.Url.Validate(); err != nil {
		return err
	}

	if err := gameInfo.ImageUrl.Validate(); err != nil {
		return err
	}

	if err := gameInfo.InitialPrice.Validate(); err != nil {
		return err
	}

	if err := gameInfo.FinalPrice.Validate(); err != nil {
		return err
	}

	if err := gameInfo.DiscountPercent.Validate(); err != nil {
		return err
	}

	return nil
}
