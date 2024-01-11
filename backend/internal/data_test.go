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

type VerifyDenseJSONCase struct {
    name string
    objs []map[string]interface{} 
    isDense bool
}

var VerifyDenseJSONCases = []VerifyDenseJSONCase {
    {
        name: "Dense Case 1",
        objs: []map[string]interface{}{
            {"f1": 5},
        },
        isDense: true,
    },
    {
        name: "Dense Case 2",
        objs: []map[string]interface{}{
            {"f1": 5},
            {"f1": 6, "f2": 12},
            {"f1": 124, "f2": 12},
        },
        isDense: true,
    },
    {
        name: "Non Dense Case 1",
        objs: []map[string]interface{}{
            {"f1": 5},
            {},
            {"f1": 124, "f2": 12},
        },
        isDense: false,
    },
    {
        name: "Non Dense Case 2",
        objs: []map[string]interface{}{
            {},
            {"f1": 124, "f2": 12},
            {},
        },
        isDense: false,
    },
    {
        name: "Non Dense Case 3",
        objs: nil,
        isDense: false,
    },
}

func TestVerifyDenseJSON(t *testing.T) {
    success := true
    for _, testCase := range VerifyDenseJSONCases {
        err := verifyDenseJSON(testCase.objs)
        isDenseResult := err == nil
        
        if isDenseResult != testCase.isDense {
            t.Logf("Case Failure: %s", testCase.name)
            success = false
        }
    }
    
    if !success {
        t.Fail()
    }
}

type StructureEqCase struct {
    name string
    s1 map[string]SDTypeID 
    s2 map[string]SDTypeID
    eq bool
}

var StructureEqCases = []StructureEqCase {
    {
        name: "Equal Case 1",
        s1: map[string]SDTypeID {"id": RealTypeID},
        s2: map[string]SDTypeID {"id": RealTypeID},
        eq: true,
    },
    {
        name: "Equal Case 2",
        s1: map[string]SDTypeID {"id": RealTypeID, "name": TextTypeID},
        s2: map[string]SDTypeID {"id": RealTypeID, "name": TextTypeID},
        eq: true,
    },
    {
        name: "Non-Equal Case 1",
        s1: map[string]SDTypeID {"id": RealTypeID},
        s2: map[string]SDTypeID {"name": TextTypeID},
        eq: false,
    },
    {
        name: "Non-Equal Case 2",
        s1: map[string]SDTypeID {"id": RealTypeID, "name": TextTypeID},
        s2: map[string]SDTypeID {"id": RealTypeID},
        eq: false,
    },
    {
        name: "Non-Equal Case 3",
        s1: map[string]SDTypeID {"id": RealTypeID},
        s2: map[string]SDTypeID {"id": RealTypeID, "name": TextTypeID},
        eq: false,
    },
}

func TestStructureEq(t *testing.T) {
    successful := true
    for _, testCase := range StructureEqCases {
        eqResult := structureEq(testCase.s1, testCase.s2)
        if eqResult != testCase.eq {
            t.Logf("Case Failure: %s", testCase.name)
            successful = false
        }
    }

    if !successful {
        t.Fail()
    }
}

type StructureDiffCase struct {
    name string
    cs map[string]SDTypeID
    rs map[string]SDTypeID
    diff map[string]SDTypeID
    shouldFail bool
}

var StructureDiffCases = []StructureDiffCase{
    {
        name: "Both Empty Case",
        cs: map[string]SDTypeID{},
        rs: map[string]SDTypeID{},
        diff: map[string]SDTypeID{},
        shouldFail: false,
    },
    {
        name: "Empty Required Structure Case",
        cs: map[string]SDTypeID{
            "col1": RealTypeID,
        },
        rs: map[string]SDTypeID{},
        diff: map[string]SDTypeID{},
        shouldFail: false,
    },
    {
        name: "Empty Current Structure Case",
        cs: map[string]SDTypeID{},
        rs: map[string]SDTypeID{
            "col1": RealTypeID,
        },
        diff: map[string]SDTypeID{
            "col1": RealTypeID,
        },
        shouldFail: false,
    },
    {
        name: "Simple Case 1",
        cs: map[string]SDTypeID{
            "col1": RealTypeID,
        },
        rs: map[string]SDTypeID{
            "col1": RealTypeID,
        },
        diff: map[string]SDTypeID{
        },
        shouldFail: false,
    },
    {
        name: "Simple Case 2",
        cs: map[string]SDTypeID{
            "col1": RealTypeID,
            "col2": RealTypeID,
        },
        rs: map[string]SDTypeID{
            "col1": RealTypeID,
            "col3": RealTypeID,
        },
        diff: map[string]SDTypeID{
            "col3": RealTypeID,
        },
        shouldFail: false,
    },
    {
        name: "Fail Case",
        cs: map[string]SDTypeID{
            "col1": RealTypeID,
        },
        rs: map[string]SDTypeID{
            "col1": TextTypeID,
        },
        shouldFail: true,
    },
}

func TestStructureDiff(t *testing.T) {
    successful := true
    for _, testCase := range StructureDiffCases {
        diff, err := structureDiff(testCase.cs, testCase.rs) 

        if err != nil && !testCase.shouldFail {
            t.Logf("Unexpected Failure %s (%s)", testCase.name, err)
            successful = false
        } else if err == nil && testCase.shouldFail {
            t.Logf("Unexpected Success %s", testCase.name)
            successful = false
        } else if !testCase.shouldFail && !structureEq(diff, testCase.diff) {
            t.Logf("Case Failure: %s %v", testCase.name, diff)
            successful = false
        }
    }

    if !successful {
        t.Fail()
    }
}


type StructureFromJSONCase struct {
    name string
    json []map[string]interface{}
    structure map[string]SDTypeID
    shouldFail bool
}

var StructureFromJSONCases = []StructureFromJSONCase{
    {
        name: "Empty Case 1",
        json: []map[string]interface{}{},
        structure: map[string]SDTypeID{},
        shouldFail: false,
    },
    {
        name: "Empty Case 2",
        json: []map[string]interface{}{{}, {}},
        structure: map[string]SDTypeID{},
        shouldFail: false,
    },
    {
        name: "Simple Case 1",
        json: []map[string]interface{}{
            {"col1": 23, "col2": "Hello"},
        },
        structure: map[string]SDTypeID{
            "col1": RealTypeID, "col2": TextTypeID,
        },
        shouldFail: false,
    },
    {
        name: "Simple Case 2",
        json: []map[string]interface{}{
            {"col1": 23},
            {"col2": "Hello"},
        },
        structure: map[string]SDTypeID{
            "col1": RealTypeID, "col2": TextTypeID,
        },
        shouldFail: false,
    },
    {
        name: "Failure Case 1",
        json: []map[string]interface{}{
            {"col1": 23},
            {"col1": "Hello"},
        },
        shouldFail: true,
    },
}


func TestStructureFromJSON(t *testing.T) {
    successful := true
    for _, testCase := range StructureFromJSONCases {
        structure, err := structureFromJSON(testCase.json)

        if err != nil && !testCase.shouldFail {
            t.Logf("Unexpected Failure %s (%s)", testCase.name, err)
            successful = false
        } else if err == nil && testCase.shouldFail {
            t.Logf("Unexpected Success %s", testCase.name)
            successful = false
        } else if !testCase.shouldFail && !structureEq(structure, testCase.structure) {
            t.Logf("Case Failure: %s %v", testCase.name, structure)
            successful = false
        }
    }

    if !successful {
        t.Fail()
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

