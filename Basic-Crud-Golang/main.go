package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Buat struct untuk data yang akan disimpan di database
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type BasicResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var db *sql.DB

func main() {
	// Inisialisasi koneksi ke database
	var err error
	db, err = sql.Open("postgres", "user=postgres dbname=basic_database password=postgre sslmode=disable port=5433")
	if err != nil {
		fmt.Print("Hello")
		log.Fatal(err)
	}
	defer db.Close()

	// Inisialisasi router
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/products", getProducts).Methods("GET")
	router.HandleFunc("/products/{id}", getProduct).Methods("GET")
	router.HandleFunc("/products", createProduct).Methods("POST")
	router.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")
	router.HandleFunc("/products", deleteProducts).Methods("DELETE")

	// Start server
	fmt.Print("Server started")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func jsonResponse(w http.ResponseWriter, products interface{}, msg string, statusCode int) {
	var response BasicResponse
	fmt.Println(products, msg)
	response.Status = statusCode
	response.Message = msg
	response.Data = products
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// Handler untuk mendapatkan semua produk
func getProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	jsonResponse(w, products, "Produk berhasil diambil", http.StatusOK)
}

// Handler untuk mendapatkan satu produk berdasarkan ID
func getProduct(w http.ResponseWriter, r *http.Request) {
	// Dapatkan ID dari variabel yang diberikan
	vars := mux.Vars(r)
	id := vars["id"]

	var product Product
	err := db.QueryRow("SELECT id, name, price FROM products WHERE id=$1", id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		jsonResponse(w, product, "Produk Error Request", http.StatusBadRequest)
		return
	}

	jsonResponse(w, product, "Produk berhasil diambil", http.StatusOK)

}

// Handler untuk menambahkan produk baru
func createProduct(w http.ResponseWriter, r *http.Request) {
	// Ambil data dari body request

	//SECTION DELIVERY
	var product Product
	var id int
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Query untuk menambahkan produk ke database
	// _, err := db.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", product.Name, product.Price)
	//SECTION REPOSITORYa
	err := db.QueryRow("INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id", product.Name, product.Price).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nid = %d\n", id)

	// SECTION USECASE
	product.ID = int(id)
	jsonResponse(w, product, "Product created successfully", http.StatusCreated)

}

// Handler untuk mengupdate produk
func updateProduct(w http.ResponseWriter, r *http.Request) {
	// Dapatkan ID dari variabel yang diberikan
	vars := mux.Vars(r)
	id := vars["id"]

	// Ambil data dari body request
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Query untuk mengupdate produk di database
	result, err := db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3", product.Name, product.Price, id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rowsAffected == 0 {
		jsonResponse(w, product, fmt.Sprintf("Tidak ada produk dengan id : %v", id), http.StatusBadRequest)
		return
	}

	jsonResponse(w, product, "Product updated successfully", http.StatusOK)
}

// Handler untuk menghapus produk
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// Dapatkan ID dari variabel yang diberikan
	vars := mux.Vars(r)
	id := vars["id"]

	// Query untuk menghapus produk dari database
	_, err := db.Exec("DELETE FROM products WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Product deleted successfully")
	jsonResponse(w, "", "Product deleted successfully", http.StatusOK)

}

// Handler untuk menghapus produk
func deleteProducts(w http.ResponseWriter, r *http.Request) {
	// Dapatkan ID dari variabel yang diberikan

	// Query untuk menghapus produk dari database
	_, err := db.Exec("DELETE FROM products ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Product deleted successfully")
	jsonResponse(w, "", "Product deleted successfully", http.StatusOK)

}
