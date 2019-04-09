package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func getMySQLDB () (db *sql.DB, err error){
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "golang"
	db,err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}