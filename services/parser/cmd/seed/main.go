package main

import (
	"context"
	repository "github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/infrastructure/sqlite"
	sqlite3Wrapper "github.com/ArtemLymarenko/steam-tg-bot/services/parser/pkg/sqlite3_wrapper"
	txmanager "github.com/ArtemLymarenko/steam-tg-bot/services/parser/pkg/tx_manager"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/seeds"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	const (
		connectionPath = "D://Development/GOLang/steam-tg-bot/services/parser/resources/sqlite/manual_test.db"
		migrationsPath = "file://D://Development/GOLang/steam-tg-bot/services/parser/resources/sqlite/migrations"
	)

	sqlite := sqlite3Wrapper.MustConnect(connectionPath, migrationsPath)
	defer sqlite.CloseConnection()

	sqlite.MustMigrateDown()
	sqlite.MustMigrateUp()

	db := sqlite.GetDbInstance()

	manager := txmanager.New(db)
	gamesRepo := repository.NewGames(db)

	const countRows = 150
	seeder := seeds.NewGameSeeder(gamesRepo, manager)
	seeder.Run(context.Background(), countRows)
}
