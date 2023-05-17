package main

import (
	"encoding/json"
	"net/http"
)

type apiError struct {
	err    string
	status int
}

func (e apiError) Error() string {
	return e.err
}

func main() {
	http.HandleFunc("/user", handleGetUserByID)
	http.ListenAndServe(":3000", nil)
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return writeJSON(w, http.StatusMethodNotAllowed, apiError{err: "invalid method", status: http.StatusMethodNotAllowed})

	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
