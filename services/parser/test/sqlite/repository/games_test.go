package repository

import (
	"context"
	"github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/domain/game"
	repository "github.com/ArtemLymarenko/steam-tg-bot/services/parser/internal/infrastructure/sqlite"
	sqlite3Wrapper "github.com/ArtemLymarenko/steam-tg-bot/services/parser/pkg/sqlite3_wrapper"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var sqlite *sqlite3Wrapper.Sqlite3

func setupRepository() game.Repository {
	return repository.NewGames(sqlite.GetDbInstance())
}

func resetDatabase() {
	sqlite.MustMigrateDown()
	sqlite.MustMigrateUp()
}

func TestMain(m *testing.M) {
	connPath := "./games.db"
	migrationPath := "file://../../../resources/sqlite/migrations"

	_, err := os.Create("./games.db")
	if err != nil {
		panic(err)
	}

	sqlite = sqlite3Wrapper.MustConnect(connPath, migrationPath)
	defer sqlite.CloseConnection()

	os.Exit(m.Run())
}

func TestCreateGame(t *testing.T) {
	resetDatabase()

	repo := setupRepository()
	ctx := context.Background()

	expectedId := game.Id(1)
	expectedName := game.Name("TestGame")
	expectedInfo := game.Info{
		GameId:          expectedId,
		Url:             "test.com",
		ImageUrl:        "test.com",
		InitialPrice:    100,
		FinalPrice:      100,
		DiscountPercent: 0,
	}

	err := repo.CreateGame(ctx, expectedId, expectedName)
	assert.NoError(t, err)

	found, err := repo.FindGame(ctx, 1)
	assert.Error(t, err, "should return error if game info not found")

	err = repo.CreateGameInfo(ctx, expectedInfo)
	assert.NoError(t, err)

	found, err = repo.FindGame(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, expectedId, found.Id, "expected and found game id should be equal")
	assert.Equal(t, expectedName, found.Name, "expected and found game name should be equal")
	assert.Equal(t, expectedInfo, found.Info, "expected and found game info should be equal")
}

func TestAdduser(t *testing.T) {
	resetDatabase()

	repo := setupRepository()
	ctx := context.Background()

	err := repo.AddUserGame(ctx, 1, 1)
	assert.Error(t, err, "should return error if user or game not found")

	expectedGames := []game.Game{
		{
			Id:   1,
			Name: "TestGame",
			Info: game.Info{
				GameId: 1,
			},
		},
		{
			Id:   2,
			Name: "TestGame2",
			Info: game.Info{
				GameId: 2,
			},
		},
	}

	for _, g := range expectedGames {
		err := repo.CreateGame(ctx, g.Id, g.Name)
		assert.NoError(t, err)
		err = repo.CreateGameInfo(ctx, g.Info)
		assert.NoError(t, err)
	}

	err = repo.AddUserGame(ctx, 1, 1)
	assert.NoError(t, err)
	err = repo.AddUserGame(ctx, 1, 2)
	assert.NoError(t, err)

	foundGames, err := repo.FindUserGames(ctx, 1)
	for i, g := range foundGames {
		assert.Equal(t, expectedGames[i].Id, g.Id, "expected and found game id should be equal")
		assert.Equal(t, expectedGames[i].Name, g.Name, "expected and found game name should be equal")
		assert.Equal(t, expectedGames[i].Info, g.Info, "expected and found game info should be equal")
	}
}

func TestSearchByName(t *testing.T) {
	resetDatabase()

	repo := setupRepository()
	ctx := context.Background()

	expectedGames := []game.Game{
		{
			Id:   1,
			Name: "Test Game",
			Info: game.Info{
				GameId: 1,
			},
		},
		{
			Id:   2,
			Name: "Test Game 2",
			Info: game.Info{
				GameId: 2,
			},
		},
	}

	for _, g := range expectedGames {
		err := repo.CreateGame(ctx, g.Id, g.Name)
		assert.NoError(t, err)
		err = repo.CreateGameInfo(ctx, g.Info)
		assert.NoError(t, err)
	}

	games, err := repo.SearchGamesByName(ctx, "Test")
	assert.NoError(t, err, "should return empty slice if no games found")
	for i, g := range games {
		assert.Equal(t, expectedGames[i].Id, g.Id, "expected and found game id should be equal")
		assert.Equal(t, expectedGames[i].Name, g.Name, "expected and found game name should be equal")
		assert.Equal(t, expectedGames[i].Info, g.Info, "expected and found game info should be equal")
	}

	games, err = repo.SearchGamesByName(ctx, "2")
	assert.NoError(t, err, "should return empty slice if no games found")
	assert.Len(t, games, 1, "should return only one game")
	assert.Equal(t, expectedGames[1].Id, games[0].Id, "expected and found game id should be equal")
	assert.Equal(t, expectedGames[1].Name, games[0].Name, "expected and found game name should be equal")
	assert.Equal(t, expectedGames[1].Info, games[0].Info, "expected and found game info should be equal")
}
