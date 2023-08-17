package cli

import (
	"FINAL_PROJECT_PHASE1/handlers"
	"bufio"
	"database/sql"
	"fmt"
)

func HandleMenu(db *sql.DB, scanner *bufio.Scanner) {
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. List Laptops")
		fmt.Println("4. Buy")
		fmt.Println("5. Edit User")
		fmt.Println("6. Delete User")
		fmt.Println("7. User Report")
		fmt.Println("8. Order Report")
		fmt.Println("9. Laptop Stock Report")
		fmt.Println("0. Exit")
		fmt.Print("Select an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			handlers.RegisterUser(db, scanner)
		case "2":
			handlers.LoginUser(db, scanner)
		case "3":
			handlers.ListLaptops(db)
		case "4":
			handlers.BuyLaptop(db, scanner)
		case "5":
			handlers.EditUser(db, scanner)
		case "6":
			handlers.DeleteUser(db, scanner)
		case "7":
			handlers.PrintUserReport(db, scanner)
		case "8":
			handlers.PrintOrderReport(db, scanner)
		case "9":
			handlers.PrintStockLaptopReport(db, scanner)
		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
