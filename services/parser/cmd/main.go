package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"parser/internal/domain"
	repository "parser/internal/infrastructure/sqlite"
	sqlitemap "parser/internal/infrastructure/sqlite/mapper"
)

func main() {
	connectionPath := "D:\\Development\\GOLang\\steam-tg-bot\\services\\parser\\resources\\sqlite\\sqlite.db"

	db, err := sql.Open("sqlite", connectionPath)
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

	tx, _ := db.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	repo := repository.New(tx)
	ctx := context.Background()
	gameInfo := domain.GameInfo{
		GameId:          2,
		ImageUrl:        "asd",
		InitialPrice:    10.2,
		FinalPrice:      10.2,
		DiscountPercent: 0,
	}
	game := domain.Game{
		Id:       2,
		Name:     "Uncharted 2",
		GameInfo: gameInfo,
	}

	err = repo.CreateGame(ctx, sqlitemap.GameEntityToCreateGameParams(game))
	if err != nil {
		log.Fatal(err)
	}

	err = repo.CreateGameInfo(ctx, sqlitemap.GameInfoEntityToCreateGameInfoParams(gameInfo))
	if err != nil {
		log.Fatal(err)
	}

	gameRow, err := repo.FindGame(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	_ = tx.Commit()

	fmt.Println(sqlitemap.FindGameRowToGame(gameRow))
}
