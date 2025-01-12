package sqlite3Wrapper

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"log"
)

type Sqlite3 struct {
	db       *sql.DB
	connPath string
}

func MustConnect(connPath string) *Sqlite3 {
	const SqliteEnableFkCmd = "PRAGMA foreign_keys = ON;"

	db, err := sql.Open("sqlite3", connPath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(SqliteEnableFkCmd)
	if err != nil {
		db.Close()
		log.Fatal(err)
	}

	return &Sqlite3{
		db:       db,
		connPath: connPath,
	}
}

func (s *Sqlite3) GetDbInstance() *sql.DB {
	return s.db
}

func (s *Sqlite3) CloseConnection() {
	if err := s.db.Close(); err != nil {
		log.Println(err)
	}
}

func (s *Sqlite3) MustMigrateUp(migrationPath string) {
	driver, err := sqlite3.WithInstance(s.db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationPath, s.connPath, driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}
}
