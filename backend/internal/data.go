package internal

import (
	sql "database/sql"
	"errors"
	"fmt"
    "strings"
)

type SDTypeID uint

// Unsigned integer identifiers for each suported type.
const (
    IntTypeID       = iota 
    TextTypeID    = iota     

    // TODO: Consider adding more.
)

// Type identifiers -> SQL type names.
var SDTypeNames = map[SDTypeID]string {
    IntTypeID:      "INTEGER",
    TextTypeID:     "TEXT",
}

var SDTypeIDs = map[string]SDTypeID{
    "INTEGER":  IntTypeID,
    "TEXT":     TextTypeID,
}

func GetTypeID(t interface{}) (SDTypeID, error) {
    switch t.(type) {
    case int:
        return IntTypeID, nil
    case string:
        return TextTypeID, nil
    }

    return 0, errors.New("Unknown Type")
}

// Given a slice of JSON objects, returns a map which maps each
// unique key name found to its corresponding type.
//
// Returns an error if an unknown type is found or the same key is mapped
// to multiple different types.
func uniqueKeys(objs []map[string]interface{}) (map[string]SDTypeID, error) {
    m := make(map[string]SDTypeID)

    for _, obj := range objs {
        for key, val := range obj {
            actualTID, err := GetTypeID(val)

            if err != nil {
                return nil, err
            }

            expectedTID, ok := m[key]
            
            if !ok {
                m[key] = actualTID
            } else if actualTID != expectedTID {
                return nil, fmt.Errorf("Type mismatch on field: %s", key)
            }
        }
    }

    return m, nil
}

func getStructureString(structure map[string]SDTypeID) string {
    var sb strings.Builder 

    i := 0
    for columnName, typeID := range structure {
        sb.WriteString(columnName)
        sb.WriteString(" ")
        sb.WriteString(SDTypeNames[typeID])
        
        if i < len(structure) - 1 {
            sb.WriteString(", ")
        }

        i++
    }

    return sb.String()
}

// Return the table structure in the form of a map which maps column
// names to data types.
//
// NOTE: This assumes the given table exists.
func getTableStructure(db *sql.DB, table string) (map[string]SDTypeID, error) {
    rows, err := db.Query("PRAGMA table_info(?)", table)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    structure := make(map[string]SDTypeID)

    var name string
    var typeName string

    for rows.Next() {
        err = rows.Scan(&name, &typeName) 

        if err != nil {
            return nil, err
        }

        tid, ok := SDTypeIDs[typeName]
        if !ok {
            return nil, fmt.Errorf("Unknown type %s.", typeName)
        }

        structure[name] = tid
    }

    return structure, nil
}

func conformTable(table string, reqColumns map[string]SDTypeID) error {
    return nil
}

func insert(db *sql.DB, table string, data []map[string]interface{}) error {
    return nil
}
