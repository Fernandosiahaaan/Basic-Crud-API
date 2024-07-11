package models

import (
	_ "github.com/mattn/go-sqlite3"
)

// struct of column table sql
type Product struct {
	ID    int    `json: "id"`
	Name  string `json : "name"`
	Price int    `json : "price"`
}

// struct of detail response
type BasicResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var MainTable string = "products"
