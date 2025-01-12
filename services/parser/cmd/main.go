package main

import (
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/app"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//connectionPath := "D:\\Development\\GOLang\\steam-tg-bot\\services\\parser\\resources\\sqlite\\sqlite.db"
	//sqlite := sqlite3Wrapper.MustConnect(connectionPath)
	//defer sqlite.CloseConnection()
	//
	//migrations := "file://services/parser/resources/sqlite/migrations"
	//sqlite.MustMigrateUp(migrations)
	//
	//db := sqlite.GetDbInstance()
	//manager := txmanager.New(db)
	//gamesRepo := repository.NewGames(db)
	//_ = service.NewGames(gamesRepo, manager)

	application := app.New(44044)
	application.Start()
}
