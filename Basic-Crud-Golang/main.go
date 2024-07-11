package main

import (
	"crud-go-sqlite/controllers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Example Crud API using SQLITE3")

	// Memuat variabel lingkungan dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// get start database
	dbName := os.Getenv("DB_NAME")
	db, err := initDB(dbName)
	if err != nil {
		log.Fatal(err)
	}

	userController := controllers.ProductController{DB: db}
	router := mux.NewRouter()
	router.HandleFunc("/products", userController.GetAllProduct).Methods("GET")
	router.HandleFunc("/products/{id}", userController.GetProductById).Methods("GET")
	router.HandleFunc("/products", userController.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", userController.UpdateProductById).Methods("PUT")
	router.HandleFunc("/products/{id}", userController.DeleteProductByID).Methods("DELETE")

	// Start server
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8000", router))
}
