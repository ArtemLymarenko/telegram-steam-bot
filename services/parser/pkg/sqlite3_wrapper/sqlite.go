package sqlite3Wrapper

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"log"
)

type Sqlite3 struct {
	db       *sql.DB
	m        *migrate.Migrate
	driver   database.Driver
	connPath string
}

func MustConnect(connPath, migrationPath string) *Sqlite3 {
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

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		db.Close()
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationPath, connPath, driver)
	if err != nil {
		db.Close()
		log.Fatal(err)
	}

	return &Sqlite3{
		db:       db,
		m:        m,
		driver:   driver,
		connPath: connPath,
	}
}

func (s *Sqlite3) GetDbInstance() *sql.DB {
	return s.db
}

func (s *Sqlite3) CloseConnection() {
	log.Println("Sqlite in closing connection...")
	if s.m != nil {
		if _, err := s.m.Close(); err != nil {
			log.Println(err)
		}
	}
	if s.driver != nil {
		if err := s.driver.Close(); err != nil {
			log.Println(err)
		}
	}
	if s.db != nil {
		if err := s.db.Close(); err != nil {
			log.Println(err)
		}
	}

	log.Println("Sqlite connection closed!")
}

func (s *Sqlite3) MustMigrateUp() {
	err := s.m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}
}

func (s *Sqlite3) MustMigrateDown() {
	err := s.m.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}
}

func (s *Sqlite3) MustMigrateDrop() {
	err := s.m.Drop()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}
}
