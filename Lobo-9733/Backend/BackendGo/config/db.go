package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error
	connStr := "user=postgres password=password dbname=instagram host=localhost port=5432 sslmode=disable"

	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexión a la base de datos exitosa")
}
