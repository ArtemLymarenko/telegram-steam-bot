package main

import (
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/app"
	sqlite3Wrapper "github.com/ArtemLymarenko/steam-tg-bot/services/parser/pkg/sqlite3_wrapper"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	const (
		connectionPath = "D:\\Development\\GOLang\\steam-tg-bot\\services\\parser\\resources\\sqlite\\sqlite.db"
		migrationsPath = "file://services/parser/resources/sqlite/migrations"
		port           = 44044
	)

	sqlite := sqlite3Wrapper.MustConnect(connectionPath)
	sqlite.MustMigrateUp(migrationsPath)

	application := app.New(port, sqlite)
	application.Start()
}
