package main

import (
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/infrastructure/external/game_parsers"
	parser_service "github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/service/parser"
	"github.com/go-co-op/gocron/v2"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
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
	factory := game_parsers.NewFactory(nil)
	steamParserApi := factory.CreateInstance(game_parsers.SteamParser)
	epicParserApi := factory.CreateInstance(game_parsers.EpicGamesParser)

	parserSvc := parser_service.New(nil)
	steamParserConfig := &parser_service.ParserConfig{
		Workers: 5,
		Parser:  steamParserApi,
	}
	epicParserConfig := &parser_service.ParserConfig{
		Workers: 5,
		Parser:  epicParserApi,
	}

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}
	defer scheduler.Shutdown()

	_, err = scheduler.NewJob(
		gocron.MonthlyJob(1,
			gocron.NewDaysOfTheMonth(-1),
			gocron.NewAtTimes(gocron.NewAtTime(24, 0, 0))),
		gocron.NewTask(
			parserSvc.RunParsersAsync,
			parserSvc.ParseGameInfoAsync, steamParserConfig, epicParserConfig,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	scheduler.Start()
}
