package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Running server ...")

	EnvLoad()
	port := os.Getenv("PORT")

	r := mux.NewRouter()
	r.HandleFunc("/ok", HealthCheck)
	r.HandleFunc("/v1/users/{userID}", User).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func EnvLoad() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

	// Preflight request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	userID := mux.Vars(r)["userID"]

	dbURL := os.Getenv("DATABASE_URL")[len("mysql://"):]
	dbPw := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbUser := "root"
	dbHost := dbURL[strings.Index(dbURL, "@")+1 : strings.LastIndex(dbURL, "/")]
	dbName := dbURL[strings.LastIndex(dbURL, "/")+1:]

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?interpolateParams=true", dbUser, dbPw, dbHost, dbName))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	row := db.QueryRow("SELECT * FROM users WHERE id = ?", userID)

	var id, age int
	var name string
	row.Scan(&id, &name, &age)

	w.Write([]byte("Hi! " + name + ""))
}
