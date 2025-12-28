package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

var schema = `
CREATE TABLE scores (
    id text,
    score INTGER NOT NULL
);`

func Connect() {
	var err error
	DB, err = sqlx.Connect("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	DB.MustExec(schema)
}
