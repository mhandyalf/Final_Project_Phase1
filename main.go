package main

import (
	"FINAL_PROJECT_PHASE1/database"
	"FINAL_PROJECT_PHASE1/handlers"
	"bufio"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to E-commerce CLI")
	fmt.Println("-------------------------")

	handlers.HandleMenu(db, scanner)
}
