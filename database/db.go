package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:vMYqpvrhJvcuJrwbhown@tcp(containers-us-west-106.railway.app:7207)/railway")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
