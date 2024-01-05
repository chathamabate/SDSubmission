package internal

import (
	sql "database/sql"
	"errors"
)

func Foo() {

}


// The user can insert a single map resulting in one row being inserted.
// Or the user can insert a list of maps resulting in multiple rows being
// inserted.
func insert(db *sql.DB, table string, data interface{}) error {
    switch v := data.(type) {

    // Single Row Case.
    case map[string]interface{}:
        break

    // Multiple Row Case.
    case []map[string]interface{}:
        break

        // In this case, let's just 

    default:
        return errors.New("Unknown data format")
    }

    
    return nil
}
