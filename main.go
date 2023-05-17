package main

import "net/http"

func main() {
	http.HandleFunc("/user", handleGetUserByID)
	http.ListenAndServe(":3000", nil)
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) {

}
