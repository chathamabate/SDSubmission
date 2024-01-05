package internal

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const (
    DEF_DB = "_db.db" 
)

func ErrorIf(t *testing.T, err error, msg string) {
    if err != nil {
        t.Error(msg)
    }
}

func TestVerifyDenseJSON(t *testing.T) {
    json := make([]map[string]interface{}, 0)
    
    if verifyDenseJSON(json) == nil {
        t.Error("Empty slice is marked dense.")
    }

    json = append(json, map[string]interface{}{
        "col1": 2,
        "col2": "text",
    })

    if verifyDenseJSON(json) != nil {
        t.Error("Dense slice is marked not dense.")
    }

    json = append(json, make(map[string]interface{}))

    if verifyDenseJSON(json) == nil {
        t.Error("Not dense slice is marked dense.")
    }
}

func TestStructureFromJSON(t *testing.T) {
    json := []map[string]interface{}{
        {
            "name": "Bob",
        },
        {
            "age": 22,
        },
        {
            "name": "Dave",
            "age": 23,
            "town": "Dover",
        },
        {
            "name": "Mark",
            "zip": 22332,
        },
    }

    acts, err := structureFromJSON(json)
    ErrorIf(t, err, "Error reading JSON")

    exps := map[string]SDTypeID {
        "name": TextTypeID,
        "age": IntTypeID,
        "town": TextTypeID,
        "zip": IntTypeID,
    }

    if !structureEq(acts, exps) {
        t.Error("Structure mismatch.")
    }
}

func prepareDB(t *testing.T) *sql.DB {
    os.Remove(DEF_DB)    
    db, err := sql.Open("sqlite3", DEF_DB)
    ErrorIf(t, err, "Error opening db.")

    return db
}

func TestStructureFromTable(t *testing.T) {
    db := prepareDB(t)
    defer db.Close()

    _, err := db.Query(`
        CREATE TABLE t1 (col1 TEXT, col2 INTEGER);
    `)

    ErrorIf(t, err, "Error creating table.")
    
    acts, err := structureFromTable(db, "t1")
    ErrorIf(t, err, "Error creating structure.")

    exps := map[string]SDTypeID {
        "col1": TextTypeID,
        "col2": IntTypeID,
    }

    if !structureEq(acts, exps) {
        t.Error("Structure mismatch.")
    }
}

