package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Function to get all product in table sql
func GetAllProduct(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select * from " + MainTable)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		products = append(products, p)
		fmt.Printf("[product] ID = %d; Name = %s; Price = %d\n", p.ID, p.Name, p.Price)
	}
	return products, nil
}

// Function to get product by id in table sql
func GetProductById(database *sql.DB, id int) (Product, error) {
	var product Product
	selectProductID := "select id, name, price from " + MainTable + " where id=$1"
	err := database.QueryRow(selectProductID, id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		log.Fatal(err)
		return product, err
	}
	return product, nil
}

// Function to create product in table sql
func CreateProduct(database *sql.DB, name string, price int) error {
	// fill data in to table users
	insertSQL := "INSERT INTO " + MainTable + " (name, price) VALUES ($1, $2)"
	_, err := database.Exec(insertSQL, name, price)
	if err != nil {
		log.Printf("[ERROR] Failed Create Product %s. Error Key = %q\n", name, err)
	}
	return err
}

// Function to update product by id in table sql
func UpdateProduct(database *sql.DB, id int, name string, price int) error {
	// fill data in to table users
	updateSQL := "UPDATE " + MainTable + " SET name=$1, price=$2 WHERE id=$3"
	result, err := database.Exec(updateSQL, name, price, id)
	if err != nil {
		log.Printf("[ERROR] Failed Update Product id %d. Error Key = %q\n", id, err)
	}
	// check result sql
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[ERROR] Failed Update Product id %d. Error Key = %q\n", id, err)
	} else if rowAffected == 0 {
		log.Printf("[ERROR] Failed Update Product id %d. Error Key = %q\n", id, err)
		return errors.New("not found id product")
	}
	return err
}

// Function to delete product by id in table sql
func DeleteProduct(database *sql.DB, id int) error {
	// fill data in to table users
	deleteSQL := "DELETE FROM " + MainTable + " WHERE id=$1"
	result, err := database.Exec(deleteSQL, id)
	if err != nil {
		log.Printf("[ERROR] Failed Delete Product id %d. Error Key = %q\n", id, err)
	}

	// check result sql
	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[ERROR] Failed Update Product id %d. Error Key = %q\n", id, err)
	} else if rowAffected == 0 {
		log.Printf("[ERROR] Failed Update Product id %d. Error Key = %q\n", id, err)
		return errors.New("Not Found ID Product")
	}
	return err
}
