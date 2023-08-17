package cli

import (
	"FINAL_PROJECT_PHASE1/handlers"
	"bufio"
	"database/sql"
	"fmt"
)

func HandleMenu(db *sql.DB, scanner *bufio.Scanner) {
	fmt.Println("Welcome to HandyComp!")

	for {
		fmt.Println("Options:")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("0. Exit")
		fmt.Print("Select an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			handlers.RegisterUser(db, scanner)
		case "2":
			loggedIn := handlers.LoginUser(db, scanner)
			if loggedIn {
				// Move to the main menu after successful login
				MainOptionsMenu(db, scanner)
			}
		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

func MainOptionsMenu(db *sql.DB, scanner *bufio.Scanner) {
	for {
		fmt.Println("Menu:")
		fmt.Println("1. List Laptops")
		fmt.Println("2. Buy")
		fmt.Println("3. Edit User")
		fmt.Println("4. Delete User")
		fmt.Println("5. User Report")
		fmt.Println("6. Order Report")
		fmt.Println("7. Laptop Stock Report")
		fmt.Println("0. Logout and Exit")
		fmt.Print("Select an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			handlers.ListLaptops(db)
		case "2":
			handlers.BuyLaptop(db, scanner)
		case "3":
			handlers.EditUser(db, scanner)
		case "4":
			handlers.DeleteUser(db, scanner)
		case "5":
			handlers.PrintUserReport(db, scanner)
		case "6":
			handlers.PrintOrderReport(db, scanner)
		case "7":
			handlers.PrintStockLaptopReport(db, scanner)
		case "0":
			fmt.Println("Logged out. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
