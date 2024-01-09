package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chathamabate/SDSubmission/backend/internal"
)


func main() {
    // Very simple logging strategy here.
    f, err := os.OpenFile("sd.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error openning log file.")
        os.Exit(1)
    }
    log.SetOutput(f)

    err = internal.RunSQLiteServer("DB.db")    

    if err != nil {
        log.Println(err)
    }
}



