package db

import (
	"database/sql"
	"log"

    _ "github.com/mattn/go-sqlite3"
)

type DB struct {
    self *sql.DB
}

const (
    XDG_DATA_HOME = "/Users/luke/.local/share"
)

func New() (*DB, error) {
    db, err := sql.Open("sqlite3", XDG_DATA_HOME + "/betterpals" + "/betterpals.db")
    if err != nil {
        return nil, err
    }
    result := &DB{db}
    result.validateTables()
    return result, nil
}

func (d *DB) validateTables() {
    _, err := d.self.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL)`)
    if err != nil {
        log.Fatal(err)
    }
}

func (d *DB) Close() {
    d.self.Close()
}
