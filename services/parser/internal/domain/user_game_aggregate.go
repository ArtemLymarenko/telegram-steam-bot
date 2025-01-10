package domain

type UserGamesAggregate struct {
	UserId int64
	Games  []*Game
}
