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
        name: "Case 1",
        objs: []map[string]interface{}{
            {"Hello": 1},
            {"Goodbye": 1},
        },
        expUks: []string{
            "Hello", "Goodbye",
        },
    },
}

func TestUniqueKeys(t *testing.T) {
    /*
    for _, uksCase := range TestUniqueKeysCases {
        actualUks := uniqueKeys(uksCase.objs)

    }*/
}
