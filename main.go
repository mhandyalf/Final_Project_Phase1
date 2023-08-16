package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/final_project_p1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to E-commerce CLI")
	fmt.Println("-------------------------")

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. List Laptops")
		fmt.Println("4. Buy")
		fmt.Println("5. Edit User")
		fmt.Println("6. Delete User")
		fmt.Println("0. Exit")
		fmt.Print("Select an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			registerUser(db, scanner)
		case "2":
			loginUser(db, scanner)
		case "3":
			listLaptops(db)
		case "4":
			buyLaptop(db, scanner)
		case "5":
			editUser(db, scanner)
		case "6":
			deleteUser(db, scanner)
		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

func registerUser(db *sql.DB, scanner *bufio.Scanner) {
	fmt.Println("Register")
	fmt.Println("--------")

	fmt.Print("Enter username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Enter email: ")
	scanner.Scan()
	email := scanner.Text()

	fmt.Print("Enter password: ")
	scanner.Scan()
	password := scanner.Text()

	// Meng-hash password sebelum menyimpannya
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	createdAt := time.Now()

	_, err = db.Exec("INSERT INTO Users (username, email, password, created_at) VALUES (?, ?, ?, ?)",
		username, email, hashedPassword, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User registered successfully!")
}

func loginUser(db *sql.DB, scanner *bufio.Scanner) {
	fmt.Println("Login")
	fmt.Println("-----")

	fmt.Print("Enter username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Enter password: ")
	scanner.Scan()
	password := scanner.Text()

	var storedPassword []byte
	var userID int
	err := db.QueryRow("SELECT user_id, password FROM Users WHERE username = ?", username).Scan(&userID, &storedPassword)
	if err != nil {
		log.Fatal(err)
	}

	// Memeriksa apakah password cocok dengan hashed password yang disimpan
	if err := bcrypt.CompareHashAndPassword(storedPassword, []byte(password)); err != nil {
		log.Fatal("Login failed: Incorrect username or password")
	}

	fmt.Println("Login successful!")
}

func listLaptops(db *sql.DB) {
	fmt.Println("List Laptops")
	fmt.Println("------------")

	rows, err := db.Query("SELECT laptop_id, brand, model, price, stock_quantity FROM Laptops")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Laptop list:")
	for rows.Next() {
		var laptopID int
		var brand, model string
		var price float64
		var stockQuantity int

		err := rows.Scan(&laptopID, &brand, &model, &price, &stockQuantity)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID: %d, Brand: %s, Model: %s, Price: %.2f, Stock: %d\n", laptopID, brand, model, price, stockQuantity)
	}
}

func buyLaptop(db *sql.DB, scanner *bufio.Scanner) {
	fmt.Println("Buy Laptop")
	fmt.Println("----------")

	fmt.Print("Enter user ID: ")
	scanner.Scan()
	userIDStr := scanner.Text()
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Fatal(err)
	}

	listLaptops(db)

	fmt.Print("Enter laptop ID to buy: ")
	scanner.Scan()
	laptopIDStr := scanner.Text()
	laptopID, err := strconv.Atoi(laptopIDStr)
	if err != nil {
		log.Fatal(err)
	}

	var laptopPrice float64
	err = db.QueryRow("SELECT price FROM Laptops WHERE laptop_id = ?", laptopID).Scan(&laptopPrice)
	if err != nil {
		log.Fatal(err)
	}

	// Assuming you have an Orders table with user_id, order_date, and total_amount columns
	orderDate := time.Now()
	_, err = db.Exec("INSERT INTO Orders (user_id, order_date, total_amount) VALUES (?, ?, ?)",
		userID, orderDate, laptopPrice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Laptop purchased successfully!")
}

func editUser(db *sql.DB, scanner *bufio.Scanner) {
	fmt.Println("Edit User")
	fmt.Println("---------")

	fmt.Print("Enter user ID: ")
	scanner.Scan()
	userIDStr := scanner.Text()
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter new full name: ")
	scanner.Scan()
	fullName := scanner.Text()

	fmt.Print("Enter new address: ")
	scanner.Scan()
	address := scanner.Text()

	fmt.Print("Enter new phone number: ")
	scanner.Scan()
	phoneNumber := scanner.Text()

	_, err = db.Exec("UPDATE User_Profiles SET full_name = ?, address = ?, phone_number = ? WHERE user_id = ?",
		fullName, address, phoneNumber, userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User information updated successfully!")
}

func deleteUser(db *sql.DB, scanner *bufio.Scanner) {
	fmt.Println("Delete User")
	fmt.Println("-----------")

	fmt.Print("Enter user ID: ")
	scanner.Scan()
	userIDStr := scanner.Text()
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM Users WHERE user_id = ?", userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User deleted successfully!")
}
