package main

import (
	"crud-go-sqlite/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var databaseName string = "Example.db"

func main() {
	fmt.Println("Example Crud API using SQLITE3")
	db, err := initDB(databaseName)
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
