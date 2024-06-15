package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	// _ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Database *sqlx.DB
}

func ConnectDB() *DB {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)
	db, err := sqlx.Connect("mysql", dns)
	if err != nil {
		log.Fatalln(err)
	}
	return &DB{Database: db}
}
