package main

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	repository "github.com/steam-tg-bot/services/parser/internal/infrastructure/sqlite"
	"github.com/steam-tg-bot/services/parser/internal/service"
	txmanager "github.com/steam-tg-bot/services/parser/pkg/tx_manager"
	"log"
)

const SqliteEnableFkCmd = "PRAGMA foreign_keys = ON;"

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

	_, err = db.Exec(SqliteEnableFkCmd)
	if err != nil {
		log.Fatal("Error enabling foreign keys: ", err)
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://services/parser/resources/sqlite/migrations", connectionPath, driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	manager := txmanager.New(db)
	gamesRepo := repository.NewGames(db)
	_ = service.NewGames(gamesRepo, manager)
}
