package internal

import (
	"testing"
)

type TestEqualStringCollectionsCase struct {
    name    string
    c1      []string 
    c2      []string
    exp     bool
}

var TestEqualStringCollectionsCases = []TestEqualStringCollectionsCase {
    {
        name: "Single Unique Element 1",
        c1: []string{"Hello"},
        c2: []string{"Hello"},
        exp: true,
    },
    {
        name: "Single Unique Element 2",
        c1: []string{"Hello"},
        c2: []string{"Hello", "Hello"},
        exp: false,
    },
    {
        name: "Simple Case 1",
        c1: []string{"Hello", "GoodBye"},
        c2: []string{"GoodBye", "Hello"},
        exp: true,
    },
    {
        name: "Simple Case 2",
        c1: []string{"Hello", "GoodBye"},
        c2: []string{"GoodBye", "Hello", "Hello"},
        exp: false,
    },
    {
        name: "Simple Case 3",
        c1: []string{"Hello", "GoodBye", "Hey", "Hey"},
        c2: []string{"Bleh", "Hello"},
        exp: false,
    },
    {
        name: "Big Case 1",
        c1: []string{"A", "B", "D", "D", "C", "A"},
        c2: []string{"B", "D", "A", "D", "C", "A"},
        exp: true,
    },
}

func TestEqualStringCollections(t *testing.T) {
    failure := false

    for _, escCase := range TestEqualStringCollectionsCases {
        res := equalStringCollections(escCase.c1, escCase.c2)

        if res != escCase.exp {
            failure = true
            t.Logf("Failure @ %s", escCase.name)
        }
    }

    if failure {
        t.Fail()
    }
}


type TestUniqueKeysCase struct {
    name    string
    objs    []map[string]interface{}
    expUks  []string
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
