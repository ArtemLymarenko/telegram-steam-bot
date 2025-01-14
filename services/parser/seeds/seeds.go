package seeds

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain/game"
	txmanager "github.com/ArtemLymarenko/steam-tg-bot/services/parser/pkg/tx_manager"
	"github.com/go-faker/faker/v4"
	"log"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

type Games struct {
	gameRepo  game.Repository
	txManager txmanager.TxManager
}

func NewGameSeeder(gameRepo game.Repository, txManager txmanager.TxManager) *Games {
	return &Games{
		gameRepo:  gameRepo,
		txManager: txManager,
	}
}

func (gs *Games) Run(ctx context.Context, gamesCount int) {
	log.Println("Seeding games...")
	gs.mustSeedGames(ctx, gamesCount)
	log.Println("Seeding games completed")
}

func (gs *Games) generateRandomGameName(length int) game.Name {
	builder := strings.Builder{}
	for range length {
		builder.WriteString(fmt.Sprintf("%s ", faker.Word()))
	}

	return game.Name(strings.Trim(builder.String(), " "))
}

func (gs *Games) mustSeedGames(ctx context.Context, gamesCount int) {
	randSeed := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range gamesCount {
		var priceSeed float64
		randGameLength := randSeed.Intn(4) + 1
		amount, _ := faker.GetPrice().Amount(reflect.ValueOf(&priceSeed).Elem())
		amountFloat64, ok := amount.(float64)
		if !ok {
			log.Fatal("failed to seed games: failed to get price amount")
		}

		err := gs.txManager.Run(ctx, &sql.TxOptions{}, func(ctx context.Context, tx *sql.Tx) error {
			txRepo := gs.gameRepo.WithTx(tx)
			err := txRepo.CreateGame(ctx, game.Id(i), gs.generateRandomGameName(randGameLength))
			if err != nil {
				return err
			}

			err = txRepo.CreateGameInfo(ctx, game.Info{
				GameId:          game.Id(i),
				Url:             "https://shared.akamai.steamstatic.com/store_item_assets/steam/apps/1172470/header.jpg?t=1734541502",
				ImageUrl:        "https://shared.akamai.steamstatic.com/store_item_assets/steam/apps/1172470/header.jpg?t=1734541502",
				InitialPrice:    game.Price(amountFloat64),
				FinalPrice:      game.Price(amountFloat64),
				DiscountPercent: 0,
			})
			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			log.Fatal("failed to seed games: ", err)
		}
	}
}
