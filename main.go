package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Running server ...")

	r := mux.NewRouter()
	r.HandleFunc("/ok", HealthCheck)
	r.HandleFunc("/", GetHome)

	log.Fatal(http.ListenAndServe(":3000", r))
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}
