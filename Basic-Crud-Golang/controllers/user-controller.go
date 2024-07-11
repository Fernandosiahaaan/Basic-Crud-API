package controllers

import (
	"crud-go-sqlite/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductController struct {
	DB *sql.DB
}

// function create format response api
func JsonResponse(w http.ResponseWriter, products interface{}, msg string, statusCode int) {
	var response models.BasicResponse
	fmt.Println(msg)
	response.Status = statusCode
	response.Message = msg
	response.Data = products
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// function handler api to get all product
func (pc *ProductController) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	product, err := models.GetAllProduct(pc.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	JsonResponse(w, product, "Success Get All Product", http.StatusOK)
}

// function handler api to get product by id
func (pc *ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	var err error

	// Get ID Product from header
	idStr := mux.Vars(r)["id"]
	product.ID, err = strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err = models.GetProductById(pc.DB, product.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	JsonResponse(w, product, fmt.Sprintf("Success Get Product ID %v", product.ID), http.StatusOK)
}

// function handler api to update product
func (pc *ProductController) UpdateProductById(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	// Get Name and Price Product from body
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get ID Product from header
	id := mux.Vars(r)["id"]
	product.ID, err = strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process Update Product
	err = models.UpdateProduct(pc.DB, product.ID, product.Name, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	JsonResponse(w, product, fmt.Sprintf("Success Update Product ID %v", product.ID), http.StatusOK)
}

// function handler api to delete product
func (pc *ProductController) DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	// Get ID Product
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = models.DeleteProduct(pc.DB, id)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		JsonResponse(w, "", err.Error(), http.StatusBadRequest)
		return
	}
	JsonResponse(w, "", fmt.Sprintf("Success Delete Product ID %v", id), http.StatusOK)
}

// function handler api to create product
func (pc *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	// Get Name and Price Product from Body
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = models.CreateProduct(pc.DB, product.Name, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	JsonResponse(w, "", fmt.Sprintf("Success Create Product Name %s; Price %v", product.Name, product.Price), http.StatusOK)
}
