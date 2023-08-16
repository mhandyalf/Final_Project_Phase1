package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/final_project_p1")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
