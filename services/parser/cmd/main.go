package main

import (
	parserService "github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/service/parser"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	//const (
	//	connectionPath = "D://Development/GOLang/steam-tg-bot/services/parser/resources/sqlite/manual_test.db"
	//	migrationsPath = "file://services/parser/resources/sqlite/migrations"
	//	port           = 44044
	//)
	//
	//sqlite := sqlite3Wrapper.MustConnect(connectionPath, migrationsPath)
	//sqlite.MustMigrateUp()
	//
	//application := app.New(port, sqlite)
	//application.Start()
	factory := parserService.NewParserFactory(nil)
	steamParser := factory.CreateInstance(parserService.SteamParserType)
	epicParser := factory.CreateInstance(parserService.EpicGamesParserType)
	service := parserService.New(nil)
	steamParserConfig := parserService.ParserConfig{
		ReadWorkers:  5,
		WriteWorkers: 5,
		Parser:       steamParser,
	}
	epicParserConfig := parserService.ParserConfig{
		ReadWorkers:  5,
		WriteWorkers: 5,
		Parser:       epicParser,
	}

	service.RunParsers(steamParserConfig, epicParserConfig)
}
