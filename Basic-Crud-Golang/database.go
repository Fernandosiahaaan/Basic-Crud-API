package main

import (
	"database/sql"

	"crud-go-sqlite/models"

	_ "github.com/mattn/go-sqlite3"
)

func initDB(NameDB string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", NameDB)
	if err != nil {
		return nil, err
	}

	err = CreateTableProduct(db) // Buat table product jika belum ada
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTableProduct(db *sql.DB) error {
	// Buat table jika belum ada
	createTableSQL := "CREATE TABLE if NOT EXISTS " + models.MainTable + "(id integer PRIMARY KEY, name text NOT NULL, price integer NOT NULL);"
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	return nil
}
