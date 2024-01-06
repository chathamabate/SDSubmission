package internal

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const (
    DEF_DB = "./_db.db" 
)

func ErrorIf(t *testing.T, err error, msg string) {
    if err != nil {
        t.Log(err)
        t.Error(msg)
    }
}

func TestVerifyDenseJSON(t *testing.T) {
    json := make([]map[string]interface{}, 0)
    
    if verifyDenseJSON(json) == nil {
        t.Error("Empty slice is marked dense")
    }

    json = append(json, map[string]interface{}{
        "col1": 2,
        "col2": "text",
    })

    if verifyDenseJSON(json) != nil {
        t.Error("Dense slice is marked not dense")
    }

    json = append(json, make(map[string]interface{}))

    if verifyDenseJSON(json) == nil {
        t.Error("Not dense slice is marked dense")
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
        "age": RealTypeID,
        "town": TextTypeID,
        "zip": RealTypeID,
    }

    if !structureEq(acts, exps) {
        t.Error("Structure mismatch")
    }
}

func prepareDB(t *testing.T) *sql.DB {
    os.Remove(DEF_DB)    
    f, err := os.Create(DEF_DB)
    ErrorIf(t, err, "Error Creating db file")
    f.Close()

    db, err := sql.Open("sqlite3", DEF_DB)
    ErrorIf(t, err, "Error opening db")

    return db
}

func TestStructureFromTable(t *testing.T) {
    db := prepareDB(t)
    defer db.Close()

    _, err := db.Exec(`
        CREATE TABLE t1 (col1 TEXT, col2 REAL);
    `)
    ErrorIf(t, err, "Error creating table") 

    acts, err := structureFromTable(db, "t1")
    ErrorIf(t, err, "Error creating structure")

    exps := map[string]SDTypeID {
        "col1": TextTypeID,
        "col2": RealTypeID,
    }

    if !structureEq(acts, exps) {
        t.Error("Structure mismatch")
    }
}

func TestConformTable(t *testing.T) {
    db := prepareDB(t)
    defer db.Close()

    rs1 := map[string]SDTypeID{
        "col1": TextTypeID,
    }

    err := conformTable(db, "t1", rs1)
    ErrorIf(t, err, "Error conforming 1") 

    rs2 := map[string]SDTypeID{
        "col1": TextTypeID,
        "col2": RealTypeID,
        "col3": TextTypeID,
    }

    err = conformTable(db, "t1", rs2)
    ErrorIf(t, err, "Error conforming 2")

    rs3 := map[string]SDTypeID{
        "col1": TextTypeID,
        "col2": RealTypeID,
        "col4": TextTypeID,
    }

    err = conformTable(db, "t1", rs3)
    ErrorIf(t, err, "Error conforming 3")

    acts, err := structureFromTable(db, "t1")
    ErrorIf(t, err, "Error getting structure")

    exps := map[string]SDTypeID{
        "col1": TextTypeID,
        "col2": RealTypeID,
        "col3": TextTypeID,
        "col4": TextTypeID,
    }

    if !structureEq(acts, exps) {
        t.Log(acts)
        t.Error("Structure mismatch")
    }
}

func TestForceInsert(t *testing.T) {
    db := prepareDB(t)
    defer db.Close()

    // First create our table.
    cs := map[string]SDTypeID{
        "name": TextTypeID,
        "age": RealTypeID,
        "zip": RealTypeID,
        "id": RealTypeID,
    }
    err := conformTable(db, "t1", cs)
    ErrorIf(t, err, "Error conforming table")

    objs := []map[string]interface{}{
        {
            "name": "Bob",
        },
        {
            "name": "Mark", 
            "age": 24,
        },
        {
            "name": "Dave", 
            "zip": 00007,
        },
    }

    rs, err := structureFromJSON(objs)
    ErrorIf(t, err, "Error derriving structure from JSON")

    err = forceInsert(db, "t1", rs, objs)
    ErrorIf(t, err, "Error inserting")

    row := db.QueryRow("SELECT COUNT(*) FROM t1;")
    var count int
    ErrorIf(t, row.Scan(&count), "Error getting count")

    if count != len(objs) {
        t.Errorf("Incorrect row count %d", count)
    }
}

func TestInsert(t *testing.T) {
    db := prepareDB(t)
    defer db.Close()

    objs1 := []map[string]interface{}{
        {
            "name": "Bob",
        },
        {
            "name": "Mark", 
            "age": 24,
        },
        {
            "name": "Dave", 
            "zip": 00007,
        },
    }

    ErrorIf(t, insert(db, "t1", objs1), "Error with first insertion")

    objs2 := []map[string]interface{}{
        {
            "id": 12,
            "name": "Alice",
        },
        {
            "id": 14,
            "name": "Josh", 
        },
    }

    ErrorIf(t, insert(db, "t1", objs2), "Error with second insertion")

    // Confirm table structure.
    
    row := db.QueryRow("SELECT COUNT(*) FROM t1;")
    var count int
    ErrorIf(t, row.Scan(&count), "Error getting count")

    if count != len(objs1) + len(objs2) {
        t.Errorf("Incorrect row count %d", count)
    }
}

func TestQuery(t *testing.T) {
    db := prepareDB(t)
    defer db.Close()

    objs := []map[string]interface{}{
        {
            "name": "Bob",
        },
        {
            "name": "Mark", 
            "age": 24,
        },
        {
            "name": "Dave", 
            "zip": 00007,
        },
    }

    ErrorIf(t, insert(db, "t1", objs), "Error with data population")

    resObjs, err := query(db, "SELECT * FROM t1;") 
    ErrorIf(t, err, "Error executing query")

    if len(resObjs) != len(objs) {
        t.Errorf("Incorrect number of result rows %d", len(resObjs))
    }

    for _, obj := range resObjs {
        if len(obj) != 3 {
            t.Errorf("Incorrect number of columns %d", len(obj))  
        }
    }
}

