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

	// models.CreateProduct(db, "tas", 45)
	// models.UpdateProduct(db, 2, "earphone", 15)
	// models.DeleteProduct(db, 6)

	// _, err = models.GetAllProduct(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	userController := controllers.ProductController{DB: db}
	router := mux.NewRouter()
	router.HandleFunc("/products", userController.GetAllProduct).Methods("GET")
	router.HandleFunc("/products", userController.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", userController.UpdateProductById).Methods("PUT")
	router.HandleFunc("/products/{id}", userController.DeleteProductByID).Methods("DELETE")

	// Start server
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8000", router))
}
