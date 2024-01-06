package internal

import (
	sql "database/sql"
	"net/http"
    "os"
    "time"
    "context"

	_ "github.com/mattn/go-sqlite3"
)


// Listens on port 3000
func RunSQLiteServer(fn string) error {
    db, err := sql.Open("sqlite3", fn)
    if err != nil {
        return err
    }
    defer db.Close()

    mux := http.NewServeMux()
    mux.Handle("/query", queryHandler{db:db})
    mux.Handle("/data", insertHandler{db:db})

    server := &http.Server{Addr: ":3000", Handler: mux}

    // Start up server in the background.
    go func() {
        server.ListenAndServe()
    }()

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)

    // Waiting for SIGINT (kill -2)
    <-stop

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
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

func (qh queryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

type insertHandler struct {
    db *sql.DB
}

func (qh insertHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}





