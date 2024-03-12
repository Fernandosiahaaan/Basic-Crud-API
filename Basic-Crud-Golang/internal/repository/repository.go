// internal/repository/repository.go

package repository

import (
	"database/sql"
	"log"
)

func InitDB(driverName string, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

// func InsertProduct(db *sql.DB, name string, price int) (int, error) {
//     // Implement insert operation
// }

// Define other repository functions...
