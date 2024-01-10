package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/chathamabate/SDSubmission/backend/internal"
)


func main() {
    lfStr := flag.String("lf", "sd.log", "Log File")
    dbfStr := flag.String("dbf", "DB.db", "Database File")
    llStr := flag.String("ll", "INFO", "Log Level (DEBUG, INFO, WARN, ERROR)")

    flag.Parse()

    if len(*lfStr) == 0 {
        fmt.Fprintln(os.Stderr, "Empty log filename given")
        os.Exit(1)
    }

    if len(*dbfStr) == 0 {
        fmt.Fprintln(os.Stderr, "Empty db filename given")
        os.Exit(1)
    }

    var ll slog.Level
    switch (*llStr) {
    case "DEBUG":
        ll = slog.LevelDebug
    case "INFO":
        ll = slog.LevelInfo
    case "WARN":
        ll = slog.LevelWarn
    case "ERROR":
        ll = slog.LevelError
    default:
        fmt.Fprintf(os.Stderr, "Unknown log level given: %s\n", *llStr)
        os.Exit(1)
    }


    // Very simple logging strategy here.
    f, err := os.OpenFile(*lfStr, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error openning log file")
        os.Exit(1)
    }

    jsonHandler := slog.NewJSONHandler(f, &slog.HandlerOptions{Level: ll})
    jsonLogger := slog.New(jsonHandler)

    slog.SetDefault(jsonLogger)

    // Entery log errors in here...
    err = internal.RunSQLiteServer(*dbfStr)    

    if err != nil {
        slog.Error(err.Error())
    }
}



