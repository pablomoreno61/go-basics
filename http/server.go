package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/**
 * Add endpoint
 * POST URL: http://localhost:8080/db
 * JSON PAYLOAD:
 * {
 *	"key": "age",
 *	"value": 35
 * }
 * Read endpoint
 * GET URL: http://localhost:8080/db/age
 * 35 as a response is expected
 */

var (
	dbValues = map[string]interface{}{}
)

type DbRequest struct {
	Key string `json:"key"`
	Value interface{} `json:"value"`
}

type DbResponse struct {
	Error  string  `json:"error"`
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[4:]

	fmt.Fprintf(w, fmt.Sprintf("%v", dbValues[key]))
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	dbRequest := &DbRequest{}

	if err := dec.Decode(dbRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Store in map
	dbValues[dbRequest.Key] = dbRequest.Value

	// Encode & return result
	w.Header().Set("Content-Type", "application/json")
	response := DbResponse{Error:""}

	enc := json.NewEncoder(w)
	if err := enc.Encode(response); err != nil {
		// Can't return error to client here
		log.Printf("can't encode %v - %s", response, err)
	}
}

func main() {
	http.HandleFunc("/db/", readHandler)
	http.HandleFunc("/db", writeHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
