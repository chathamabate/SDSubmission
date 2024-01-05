package internal

import (
	"testing"
)

type TestUniqueKeysSuccessCase struct {
    name    string
    objs    []map[string]interface{}
    expUks  map[string]uint
}

var TestUniqueKeysCases = []TestUniqueKeysCase {
    {
        name: "Simple Case 1",
        objs: []map[string]interface{}{
            {"Hello": 1},
            {"Goodbye": 1},
        },
        expUks: []string{
            "Hello", "Goodbye",
        },
    },
    {
        name: "Simple Case 2",
        objs: []map[string]interface{}{
            {"Hello": 1, "Goodbye": 1},
            {"Goodbye": 1},
        },
        expUks: []string{
            "Hello", "Goodbye",
        },
    },
    {
        name: "Simple Case 3",
        objs: []map[string]interface{}{
            {"A": 1, "B": 1},
            {"C": 1, "B": 1},
            {"D": 1, "B": 1},
        },
        expUks: []string{
            "A", "B", "C", "D",
        },
    },
    {
        name: "Simple Case 4",
        objs: []map[string]interface{}{
            {"A": 1, "B": 1, "C": 1},
        },
        expUks: []string{
            "A", "B", "C",
        },
    },
}

func TestUniqueKeys(t *testing.T) {
    failure := false
    for _, uksCase := range TestUniqueKeysCases {
        actualUks := uniqueKeys(uksCase.objs)
        
        if !equalStringCollections(actualUks, uksCase.expUks) {
            failure = true

            t.Logf("Failure @ %s, actual: %v", 
                uksCase.name, actualUks)
        }
    }

    if failure {
        t.Fail()
    }
}
