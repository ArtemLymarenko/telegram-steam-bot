package game

type UserId int64

type UserGames struct {
	UserId UserId
	Games  []Game
}
