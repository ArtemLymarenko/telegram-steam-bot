package domain

type Game struct {
	Id       int64
	Name     string
	GameInfo GameInfo
}

func (game Game) Validate() bool {
	if game.Id < 0 {
		return false
	}

	if game.Name == "" {
		return false
	}

	if !game.GameInfo.Validate() {
		return false
	}

	return true
}
