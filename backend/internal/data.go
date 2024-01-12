package internal

import (
	sql "database/sql"
	"errors"
	"fmt"
	"log/slog"
	"reflect"
	"strconv"
	"strings"
)

type SDTypeID uint

// Unsigned integer identifiers for each suported type.
const (
	RealTypeID = iota
	TextTypeID = iota

	// TODO: Consider adding more.
)

// Type identifiers -> SQL type names.
var SDTypeNames = map[SDTypeID]string{
	RealTypeID: "REAL",
	TextTypeID: "TEXT",
}

var SDTypeIDs = map[string]SDTypeID{
	"REAL": RealTypeID,
	"TEXT": TextTypeID,
}

// Type identifiers -> SQL Default values as strings.
var SDTypeDefaults = map[SDTypeID]string{
	RealTypeID: "0.0",
	TextTypeID: "NULL",
}

func SDGetTypeID(v interface{}) (SDTypeID, error) {
	switch v.(type) {
	case float64, int:
		return RealTypeID, nil
	case string:
		return TextTypeID, nil
	}

	return 0, errors.New("Unknown type")
}

func SDValToString(v interface{}) (string, error) {
	switch val := v.(type) {
	case int:
		return strconv.Itoa(val), nil
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64), nil
	case string:
		return "\"" + val + "\"", nil
	}

	return "", errors.New("Unknown type")
}

// This function confirms that the given slice of objects is non-empty
// AND contains no empty objects.
func verifyDenseJSON(objs []map[string]interface{}) error {
	if len(objs) == 0 {
		return errors.New("Empty object list")
	}

	for _, m := range objs {
		if len(m) == 0 {
			return errors.New("Empty object found in list")
		}
	}

	return nil
}

// NOTE: For the rest of this file we use "structure" to
// refer to the columns and types of a table.
// These are stored as a map[string]SDTypeID

func structureEq(s1 map[string]SDTypeID, s2 map[string]SDTypeID) bool {
	for k1, v1 := range s1 {
		v2, ok := s2[k1]

		if !ok || v1 != v2 {
			return false
		}
	}

	for k2, v2 := range s2 {
		v1, ok := s1[k2]

		if !ok || v1 != v2 {
			return false
		}
	}

	return true
}

// Given a current structure (cs) and a required structure (rs) returns
// a structure containing all columns in rs which are not in cs.
//
// NOTE: rs contains a column which is in cs, but maps to a different type,
// an error is returned.
func structureDiff(cs map[string]SDTypeID, rs map[string]SDTypeID) (map[string]SDTypeID, error) {
	diff := make(map[string]SDTypeID)

	for reqName, reqType := range rs {
		currType, ok := cs[reqName]

		if !ok {
			diff[reqName] = reqType
		} else if reqType != currType {
			return nil, fmt.Errorf("Type mismatch on column %s", reqName)
		}
	}

	return diff, nil
}

func structureString(s map[string]SDTypeID) string {
	var sb strings.Builder

	i := 0
	for columnName, typeID := range s {
		sb.WriteString(columnName)
		sb.WriteString(" ")
		sb.WriteString(SDTypeNames[typeID])

		if i < len(s)-1 {
			sb.WriteString(", ")
		}

		i++
	}

	return sb.String()
}

// Given a slice of JSON objects, returns a map which maps each
// unique key name found to its corresponding type.
//
// Returns an error if an unknown type is found or the same key is mapped
// to multiple different types.
func structureFromJSON(objs []map[string]interface{}) (map[string]SDTypeID, error) {
	s := make(map[string]SDTypeID)

	for _, obj := range objs {
		for key, val := range obj {
			actualTID, err := SDGetTypeID(val)

			if err != nil {
				return nil, err
			}

			expectedTID, ok := s[key]

			if !ok {
				s[key] = actualTID
			} else if actualTID != expectedTID {
				return nil, fmt.Errorf("Type mismatch on field: %s", key)
			}
		}
	}

	return s, nil
}

// Return the table structure in the form of a map which maps column
// names to data types.
//
// NOTE: This assumes the given table exists.
func structureFromTable(db *sql.DB, table string) (map[string]SDTypeID, error) {
	rows, err := db.Query("PRAGMA table_info(" + table + ");")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	s := make(map[string]SDTypeID)

	var name string
	var typeName string

	var dummyNum int
	var dummyPtr *interface{}

	for rows.Next() {
		err = rows.Scan(&dummyNum, &name, &typeName,
			&dummyNum, &dummyPtr, &dummyNum)

		if err != nil {
			return nil, err
		}

		tid, ok := SDTypeIDs[typeName]
		if !ok {
			return nil, fmt.Errorf("Unknown type %s", typeName)
		}

		s[name] = tid
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return s, nil
}




// This function will create the given table if it doesn't exist.
// Otherwise, it will alter the table if needed to conform to
// the required structure.
func conformTable(db *sql.DB, table string, rs map[string]SDTypeID) error {
	row := db.QueryRow(`
        SELECT COUNT(*)
        FROM sqlite_master
        WHERE type="table" AND name="` + table + "\";")

	var count int
	err := row.Scan(&count)

	if err != nil {
		return err
	}

	// Create a new table if needed.
	if count == 0 {
		_, err := db.Exec("CREATE TABLE " + table +
			"(" + structureString(rs) + ");")

		return err
	}

	// Otherwise, alter the current table if needed.
	cs, err := structureFromTable(db, table)
	if err != nil {
		return err
	}

	ds, err := structureDiff(cs, rs)
	if err != nil {
		return err
	}

	// NOTE: due to race conditions, I will NOT be performing error checking
	// after each column addition.
	//
	// It is possible that at the time of this addition, the required column has already
	// been added. In this case, adding the column would result in an error even
	// though nothing would go wrong after this point.
	//
	// If there is an ellusive SQL error at this step, it will slip by unnoticed.
	// If the user attempts to send new columns with incosistent data types, an error
	// might be missed as well.
	//
	// In either case, nothing will break, the user will simply get a confusing SQL
	// error message later on.
	//
	// TODO: With more time, this should be addressed in a better way.

	for colName, colType := range ds {
		_, err = db.Exec("ALTER TABLE " + table +
			" ADD COLUMN " + colName + " " + SDTypeNames[colType] + ";")

       if err != nil {
            slog.Debug("Error Altering Table", "description", err.Error())
       }
	}

	return nil
}

// Convert a JSON object of primitive values into a string ready for
// SQL insertions.
//
// NOTE: order gaurantees the ordering of columns in the resulting
// string.
func objString(rs map[string]SDTypeID, obj map[string]interface{}, order map[string]int) string {
	colStrings := make([]string, len(rs))
	i := 0

	for colName, colType := range rs {
		var err error
		var strRep string

		val, ok := obj[colName]

		if ok {
			strRep, err = SDValToString(val)
		}

		place := order[colName]

		// If there is an error creating a string representation
		// of our field, OR our object doesn't contain a value
		// for said field, we write the default string value for that
		// type instead.
		if err != nil || !ok {
			colStrings[place] = SDTypeDefaults[colType]
		} else {
			colStrings[place] = strRep
		}

		i++
	}

	return strings.Join(colStrings, ", ")
}

// This function constructs and performs the actual insert query.
// It should only be called after necessary checks and alterations
// have been done on the given table.
func forceInsert(db *sql.DB, table string, rs map[string]SDTypeID, objs []map[string]interface{}) error {
	// First create an arbitrary ordering.
	i := 0
	order := make(map[string]int)
	header := make([]string, len(rs))
	for colName := range rs {
		order[colName] = i
		header[i] = colName
		i++
	}

	headerString := strings.Join(header, ", ")

	objStrings := make([]string, len(objs))
	for i, obj := range objs {
		objStrings[i] = "(" + objString(rs, obj, order) + ")"
	}
	objsString := strings.Join(objStrings, ", ")

	query := "INSERT INTO " + table + "(" + headerString + ") VALUES " +
		objsString + ";"

	_, err := db.Exec(query)
	return err
}

// Insert Logic Flow.
//
// 1) We get a list of JSON objects from the user.
// 2) Translate the list of objects into just its aggregate structure.
// 3) Perform as needed table alterations:
//      a) If the table does not exist, create a new table with
//         the given aggregate structure.
//      b) If the table does exist, compare its columns to that
//         of the aggregate structure. Add new columns if needed,
//         report an error if the given data's structure mismatches
//         that of the prexisting columns.
// 4) Finally, insert data into the table.

func insert(db *sql.DB, table string, objs []map[string]interface{}) error {
	err := verifyDenseJSON(objs)
	if err != nil {
		return err
	}

	rs, err := structureFromJSON(objs)
	if err != nil {
		return err
	}

	err = conformTable(db, table, rs)
	if err != nil {
		return err
	}

	return forceInsert(db, table, rs, objs)
}

func query(db *sql.DB, q string) ([]map[string]interface{}, error) {
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

    colTypes, err := rows.ColumnTypes()
    if err != nil {
        return nil, err
    }

	colNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// NOTE: This approach I took almost entirely from Stackoverflow.

    // used for allocation & dereferencing
    rowValues := make([]reflect.Value, len(colTypes))
    for i, ct := range colTypes {
        rowValues[i] = reflect.New(reflect.PtrTo(ct.ScanType()))
    }

	numCols := len(colNames)
	finalRows := make([]map[string]interface{}, 0)

    for rows.Next() {
        // initially will hold pointers for Scan, after scanning the
        // pointers will be dereferenced so that the slice holds actual values
        rowResult := make([]interface{}, numCols)
        for i, rv := range rowValues {
            rowResult[i] = rv.Interface()
        }

        // scan each column value into the corresponding **T value
        err = rows.Scan(rowResult...)
        if err != nil {
            return nil, err
        }

        finalRow := make(map[string]interface{})

        for i, rv := range rowValues {
            v := rv.Elem()

            if v.IsNil() {
                finalRow[colNames[i]] = nil
                continue
            }

            switch v := rv.Elem().Interface().(type) {
            case *sql.NullFloat64:
                finalRow[colNames[i]] = v.Float64
                break
            case *sql.NullString:
                finalRow[colNames[i]] = v.String
                break
            default:
                finalRow[colNames[i]] = v
                break
            }
        }
        
        finalRows = append(finalRows, finalRow)
    }

	return finalRows, nil
}
