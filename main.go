package main

import "net/http"

func main() {
	http.HandleFunc("/user", handleGetUserByID)
	http.Listem
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) {

}
