package dto

type Game struct {
	Id              int64
	Name            string
	Url             string
	ImageUrl        string
	InitialPrice    float64
	FinalPrice      float64
	DiscountPercent float64
}

type GetUserGamesRequest struct {
	UserId int64
}

type GetUserGamesResponse struct {
	Games []Game
}

type AddUserGameRequest struct {
	UserId int64
	GameId int64
}

type AddUserGameResponse struct {
	Success bool
}

type SearchGameRequest struct {
	Name string
}

type SearchGameResponse struct {
	Games []Game
}
