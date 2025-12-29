package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

var schema = `
CREATE TABLE scores (
    id text,
    score INTGER NOT NULL,
	created DATETIME NOT NULL DEFAULT(datetime('now'))
);`

func Connect() {
	var err error
	DB, err = sqlx.Connect("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	ids := make([]string, 0)
	err = DB.Select(&ids, "SELECT DISTINCT id FROM scores") //dummy querry
	if err != nil {
		DB.MustExec(schema)
	}
}
