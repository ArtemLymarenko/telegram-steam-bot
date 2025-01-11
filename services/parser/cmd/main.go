package main

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"log"
	repository "parser/internal/infrastructure/sqlite"
	"parser/internal/service"
	txmanager "parser/pkg/tx_manager"
)

func main() {
	connectionPath := "D:\\Development\\GOLang\\steam-tg-bot\\services\\parser\\resources\\sqlite\\sqlite.db"

	db, err := sql.Open("sqlite3", connectionPath)
	if err != nil {
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			return
		}
	}()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://resources/migrations", connectionPath, driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	manager := txmanager.New(db)
	gamesRepo := repository.NewGames(db)
	gamesSvc := service.NewGames(gamesRepo, manager)
	err = gamesSvc.TxTest()
	if err != nil {
		log.Fatal(err)
	}
}
