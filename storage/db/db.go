package db

import (
	"database/sql"
	"fmt"
)

func OpenDB(path string) (*sql.DB, error){
    db, err := sql.Open("sqlite3", fmt.Sprintf("%s/betterpals.db", path))
    if err != nil {
        return nil, err
    }
    return db, nil
}
