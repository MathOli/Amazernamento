package db

import (
	"database/sql"

	"loja/check"

	_ "github.com/lib/pq"
)

func InitBd() *sql.DB {
	conexao := "user=postgres password=123casasoul dbname=alura_loja host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	check.Check(err)
	return db
}
