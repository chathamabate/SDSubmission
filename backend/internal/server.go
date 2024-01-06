package internal

import (
	"context"
	sql "database/sql"
	"encoding/json"
	"errors"
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

    log.Println("Starting server on port 3000")
    // Start up server in the background.
    go func() {
        server.ListenAndServe()
    }()

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)

    // Waiting for SIGINT (kill -2)
    <-stop

    log.Println("Shutting down server")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()  // Used for cleaning up above context.

    if err := server.Shutdown(ctx); err != nil {
        return err
    }

    return nil
}

// NOTE: 
// Both endpoints below will return a JSON object 
// with structure:
// {
//      message: string
//      data: any
// }
//
// The message field can be used arbitrarily.
// The intention is for it to hold error information
// when needed.

func writeError(w http.ResponseWriter, desc string, err error) {
    emsg := "(" + desc + ")"

    if err != nil {
        emsg += " " + err.Error()
        log.Println(emsg)
    }

    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": emsg,
        "data": nil,
    })
}
   
// These functions create closures around the given 
// thread-safe database handle.

type queryHandler struct {
    db *sql.DB
}

// Expectes URL argument q which maps to query string.
func (qh queryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    args, err := url.ParseQuery(r.URL.RawQuery)

    if err != nil {
        writeError(w, "Error parsing query from URL", err)
        return 
    }

    q, ok := args["q"]
    if !ok {
        writeError(w, "Query not provided", nil)
        return 
    }

    objs, err := query(qh.db, q[0])
    if err != nil {
        writeError(w, "Failed query", err)
        return 
    }
    
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "",
        "data": objs,
    })
}

type insertHandler struct {
    db *sql.DB
}

func (ih insertHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    args, err := url.ParseQuery(r.URL.RawQuery)

    if err != nil {
        writeError(w, "Error parsing query", err)
        return 
    }

    t, ok := args["table"]
    if !ok {
        writeError(w, "Table not provided", nil)
        return 
    }

    var reqObj interface{}
    var reqSlice []map[string]interface{}

    err = json.NewDecoder(r.Body).Decode(&reqObj)

    if err != nil {
        writeError(w, "Error decoding request body", err)
        return 
    }

    switch v := reqObj.(type) {
    case map[string]interface{}:
        reqSlice = []map[string]interface{}{v}
        break
    case []map[string]interface{}:
        reqSlice = v
        break
    default:
        writeError(w, "Unexpected request body type", nil) 
        return
    }

    err = insert(ih.db, t[0], reqSlice)
    if err != nil {
        writeError(w, "Data insertion error", err)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "",
        "data": nil,
    })
}





