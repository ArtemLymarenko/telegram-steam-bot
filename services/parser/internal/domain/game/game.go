package game

import (
	"errors"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain"
)

type Id int64

func (id Id) Validate() domain.ValidationError {
	if id < 0 {
		return errors.New("id must be greater than 0")
	}
	return nil
}

type Name string

func (name Name) Validate() domain.ValidationError {
	if name == "" {
		return errors.New("name must not be empty")
	}
	return nil
}

type Game struct {
	Id   Id
	Name Name
	Info Info
}

func (game Game) Validate() domain.ValidationError {
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
