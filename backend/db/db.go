package db

import (
	"database/sql"
	"gofy/config"

	"github.com/golang-migrate/migrate/v4"
	sqlite_migrate "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init() {
	c := config.GetConfig()
	var err error
	db, err = sql.Open("sqlite3", c.GetString("db.path"))
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}

	if err = ensureSchema(db); err != nil {
		panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}

func ensureSchema(db *sql.DB) error {
	driver, err := sqlite_migrate.WithInstance(db, &sqlite_migrate.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./db/migrations",
		"sqlite3", driver)
	if err != nil {
		return err
	}

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
