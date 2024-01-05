package internal

import (
	sql "database/sql"
)

// This function confirms that the given two slices of strings 
// contain equal values regardless of order.
func equalStringCollections(c1 []string, c2 []string) bool {
    // Count the frequencies of every string in c2
    c2Freqs := make(map[string]int)
    for _, v := range c2 {
        if _, ok := c2Freqs[v]; ok {
            c2Freqs[v]++
        } else {
            c2Freqs[v] = 1
        }
    }

    // Confirm all elements in c1 are in c2
    for _, v := range c1 {
        freq, ok := c2Freqs[v] 

        if !ok || freq == 0 {
            return false
        }

        c2Freqs[v]--
    }

    // Confirm there are no elements in c2 which
    // are not in c1
    for _, freq := range c2Freqs {
        if freq > 0 {
            return false
        }
    }

    return true
}

// Given a slice of JSON objects, returns a slice containing
// every unique key name which appears.
// (These unique keys will be mapped to SQL table columns) 
func uniqueKeys(objs []map[string]interface{}) []string {
    m := make(map[string]bool)

    for _, obj := range objs {
        for key := range obj {
            if _, ok := m[key]; !ok {
                m[key] = true
            }
        }
    }

    // Translate m into a slice of just its keys.
    keys := make([]string, len(m))

    i := 0
    for key := range m {
        keys[i] = key
        i++
    }

    return keys;
}

func confirmTable(table string, data []map[string]interface{}) ([]string, error) {
    return nil, nil
}

func insert(db *sql.DB, table string, data []map[string]interface{}) error {
    return nil
}
