package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConectaBD(search_path string) *sql.DB {
	conexao := "user=postgres dbname=HareInteractCRM password=HareWare@2024 host=localhost sslmode=disable search_path=" + search_path

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
