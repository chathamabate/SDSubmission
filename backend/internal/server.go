package internal

import (
	"context"
	sql "database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Listens on port 3000.
// Waits for SIGKILL to stop and gracefully shut down.
func RunSQLiteServer(fn string) error {
    _, err := os.Stat(fn)
    if errors.Is(err, os.ErrNotExist) {
        log.Printf("Creating new DB %s.\n", fn)
        f, err := os.Create(fn)
        if err != nil {
            return err
        }
        f.Close()
    } else if err != nil {
        return err  // funky error running stat?
    }

    log.Printf("Openning DB %s.\n", fn)

    db, err := sql.Open("sqlite3", fn)
    if err != nil {
        return err
    }
    defer db.Close()

    mux := http.NewServeMux()
    mux.Handle("/query", queryHandler{db:db})
    mux.Handle("/data", insertHandler{db:db})

    server := &http.Server{Addr: ":3000", Handler: mux}

    log.Println("Starting server on port 3000.")
    // Start up server in the background.
    go func() {
        server.ListenAndServe()
    }()

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)

    // Waiting for SIGINT (kill -2)
    <-stop

    log.Println("Shutting down server.")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()  // Used for cleaning up above context.

    if err := server.Shutdown(ctx); err != nil {
        return err
    }

    return nil
}

// These functions create closures around the given 
// thread-safe database handle.

type queryHandler struct {
    db *sql.DB
}

// Expectes URL argument q which maps to query string.
func (qh queryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    args, err := url.ParseQuery(r.URL.RawQuery)

    if err != nil {
        log.Println("Error parsing query.")
        return // TODO write error.
    }

    q, ok := args["q"]
    if !ok {
        log.Println("Query not provided.")
        return // TODO write error.
    }

    objs, err := query(qh.db, q[0])
    if err != nil {
        log.Println(err)
        return // TODO
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(objs) 
}

type insertHandler struct {
    db *sql.DB
}

func (qh insertHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from insert endpoint.\n")
}





