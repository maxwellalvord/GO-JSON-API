package main

import (
	"encoding/json"
	"net/http"
)

type apiError struct {
	Err    string
	Status int
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, apiError{Err: "internal server"})
		}
	}
}

func (e apiError) Error() string {
	return e.Err
}

func main() {
	http.HandleFunc("/user", makeHTTPHandler(handleGetUserByID))
	http.ListenAndServe(":3000", nil)
}

type User struct {
	id int
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return apiError{Err: "invalid method", Status: http.StatusMethodNotAllowed}
	}

	return writeJSON(w, http.StatusOK, User{})
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
