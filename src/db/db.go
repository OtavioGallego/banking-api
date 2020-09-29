package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL Driver
)

// Connect opens a MySQL connection and returns it
func Connect() (*sql.DB, error) {
	var stringConexao = fmt.Sprintf("root:root@tcp(db:3306)/pismo")
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
