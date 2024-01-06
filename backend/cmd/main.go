package main

import (
	"log"

	"github.com/chathamabate/SDSubmission/backend/internal"
)


func main() {
    err := internal.RunSQLiteServer("testDB.db")    
    if err != nil {
        log.Println(err)
    }
}



