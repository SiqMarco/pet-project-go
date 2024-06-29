package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // Driver do banco de dados
	"petPorject/api/src/config"
)

func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StingDbConnection)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil

}
