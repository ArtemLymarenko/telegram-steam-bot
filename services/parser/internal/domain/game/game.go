package game

import (
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain"
)

type Id int64

func (id Id) Validate() error {
	if id < 0 {
		return domain.ValidationError{
			Err: "id must be greater than 0",
		}
	}
	return nil
}

type Name string

func (name Name) Validate() error {
	if name == "" {
		return domain.ValidationError{
			Err: "name must not be empty",
		}
	}
	return nil
}

type Game struct {
	Id   Id
	Name Name
	Info Info
}

func (game Game) Validate() error {
	if err := game.Id.Validate(); err != nil {
		return err
	}

	if err := game.Name.Validate(); err != nil {
		return err
	}

	if err := game.Info.Validate(); err != nil {
		return err
	}

	return nil
}
