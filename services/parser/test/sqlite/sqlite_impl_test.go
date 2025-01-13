package sqlite

import (
	sqlite3Wrapper "github.com/ArtemLymarenko/steam-tg-bot/services/parser/pkg/sqlite3_wrapper"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	connPath      = "./test.db"
	migrationPath = "file://../../resources/sqlite/migrations"
)

func TestSqliteConnectAndMigrate(t *testing.T) {
	assert.NotPanics(t, func() {
		sqlite := sqlite3Wrapper.MustConnect(connPath, migrationPath)
		defer sqlite.CloseConnection()
		sqlite.MustMigrateUp()
		sqlite.MustMigrateDown()
		sqlite.MustMigrateDrop()
	})
}
