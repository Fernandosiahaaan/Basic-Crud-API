package main

import (
	"basic-crud-golang/internal/repository"
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {

	// initialize database
	db, _ = repository.InitDB("postgres", "user=postgres dbname=tutorial_project password=postgre sslmode=disable port=5433")
	defer db.Close()

	// Inisialisasi router
	router := mux.NewRouter()

	// // Routes
	// router.HandleFunc("/products", getProducts).Methods("GET")
	// router.HandleFunc("/products/{id}", getProduct).Methods("GET")
	// router.HandleFunc("/products", createProduct).Methods("POST")
	// router.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	// router.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")
	// router.HandleFunc("/products", deleteProducts).Methods("DELETE")

	// // Start server
	// fmt.Print("Server started")
	// log.Fatal(http.ListenAndServe(":8000", router))

}
