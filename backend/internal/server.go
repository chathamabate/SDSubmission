package internal

import (
    "net/http"
	sql "database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func RunSQLiteServer(fn string) error {
    db, err := sql.Open("sqlite3", fn)

    // Ok, so we need a result to JSON. 
    // We also need a 

    if err != nil {
        return err
    }

    defer db.Close()

    return nil
}

// These functions create closures around the given 
// thread-safe database handle.

func queryHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
    return func(res http.ResponseWriter, req *http.Request) {
        // We shall see soon enough how to do this...
    }
}

func insertHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
    return func(res http.ResponseWriter, req *http.Request) {
    }
}




