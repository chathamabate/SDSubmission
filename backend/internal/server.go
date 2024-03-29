package internal

import (
	"context"
	sql "database/sql"
	"encoding/json"
	"errors"
	"log/slog"
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
        slog.Info("Creating new DB", "db", fn)

		f, err := os.Create(fn)
		if err != nil {
            slog.Error("Couldn't create DB", "db", fn, "error", err)
			return err
		}

		f.Close()
	} else if err != nil {
        slog.Error("Couldn't create DB", "db", fn, "error", err)
		return err // funky error running stat?
	}

    slog.Info("Openning DB", "db", fn)

	db, err := sql.Open("sqlite3", fn)
	if err != nil {
        slog.Error("Couldn't open DB", "db", fn, "error", err)
		return err
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.Handle("/query", queryHandler{db: db})
	mux.Handle("/data", insertHandler{db: db})

	server := &http.Server{Addr: ":3000", Handler: mux}

    slog.Info("Starting server", "port", 3000)

	// Start up server in the background.
	go func() {
		server.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (kill -2)
	<-stop
    
    slog.Info("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Used for cleaning up above context.

	if err := server.Shutdown(ctx); err != nil {
        slog.Error("Couldn't shut down server", "error", err)
		return err
	}

    slog.Info("Server shutdown successfully")

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

func WriteError(w http.ResponseWriter, desc string, err error) {
    emsg := desc

	if err != nil {
		emsg += " (" + err.Error() + ")" 
        slog.Debug("User request error", "description", desc, "error", err)
	} else {
        slog.Debug("User request error", "description", desc)
    }

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": emsg,
		"data":    nil,
	})
}

// These functions create closures around the given
// thread-safe database handle.

type queryHandler struct {
	db *sql.DB
}

// Expectes URL argument q which maps to query string.
func (qh queryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: THIS IS NOT SUITABLE FOR PRODUCTION.
	// NOTE: This is to TEMPORARILY bypass CORS Errors (Potentially very dangerous)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token")

	w.Header().Set("Content-Type", "application/json")

	args, err := url.ParseQuery(r.URL.RawQuery)

	if err != nil {
		WriteError(w, "Error parsing query from URL", err)
		return
	}

	q, ok := args["q"]
	if !ok {
		WriteError(w, "Query not provided", nil)
		return
	}

	objs, err := query(qh.db, q[0])
	if err != nil {
		WriteError(w, "Failed query", err)
		return
	}

    slog.Debug("Query success", "query", q[0])

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "",
		"data":    objs,
	})
}

type insertHandler struct {
	db *sql.DB
}

func (ih insertHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: THIS IS NOT SUITABLE FOR PRODUCTION.
	// NOTE: This is to TEMPORARILY bypass CORS Errors (Potentially very dangerous)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token")
	w.Header().Set("Content-Type", "application/json")

	args, err := url.ParseQuery(r.URL.RawQuery)

	if err != nil {
		WriteError(w, "Error parsing query", err)
		return
	}

	t, ok := args["table"]
	if !ok {
		WriteError(w, "Table not provided", nil)
		return
	}

	var reqObj interface{}
	var reqSlice []map[string]interface{} = nil

	err = json.NewDecoder(r.Body).Decode(&reqObj)

	if err != nil {
		WriteError(w, "Error decoding request body", err)
		return
	}

	switch v := reqObj.(type) {
	case map[string]interface{}:
		reqSlice = []map[string]interface{}{v}
		break
	case []interface{}:
		reqSlice = make([]map[string]interface{}, len(v))
		for i, entry := range v {
			obj, ok := entry.(map[string]interface{})
			if !ok {
				WriteError(w, "Unexpected request body type", nil)
			}
			reqSlice[i] = obj
		}
		break
	default:
		WriteError(w, "Unexpected request body type", nil)
		return
	}

	err = insert(ih.db, t[0], reqSlice)
	if err != nil {
		WriteError(w, "Data insertion error", err)
		return
	}

    slog.Debug("Insert success", "table", t[0], "rows", len(reqSlice))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "",
		"data":    nil,
	})
}
