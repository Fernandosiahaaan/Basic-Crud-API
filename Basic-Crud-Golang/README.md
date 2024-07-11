# Basic-Crud-Golang

## Description

This Repo describe how to create basic CRUD (Create, Read, Update, and Delete) in API flow in Golang using Gorilla Mux Framework. This database will using SQL lite.

## Features

- Create/Add Products
- Get All Product
- Get Product based ID
- Update Products
- Delete Products

## Runing Project

Please Clone the project in your Local Directory

```
git clone https://github.com/Fernandosiahaaan/Basic-Crud-API.git
cd Basic-Crud-Golang
```

Before start project please add package installation for this project.

```
go mod tidy 

or

go get github.com/lib/pq	     	    # install module
go get github.com/gorilla/mux		    # install module
go get github.com/mattn/go-sqlite3		# install module
```

After already installed packages so you can run the project.

```
go run main.go

or 

go build 
./crud-go-sqlite
```

## File Tree

```
.
├── Basic-Crud-Golang             # Example Crud in Go/Golang
│   ├── controllers               # folder of controller software (handler api)
│   ├── models                    # folder of model business (query sql)
│   ├── crud-sql-lite             # binary of project
│   ├── example.db                # db of this project (using sqlite3)
│   ├── go.mod                    # init of the project
│   ├── go.sum                    # List of packages/library
│   ├── main.go                   # main of the project
│   ├── README.md                 # documentation of the project
├── Basic-Crud-Python             # Example Crud in Python
└── README.md                     # Detail of this repo

```
