package domain

type UserGamesAggregate struct {
	UserId int
	Games  []*Game
}
