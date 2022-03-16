package main

import (
	"fmt"
	// "gorm.io/gorm"
)

// var (
// 	DB *gorm.DB
// )

func main() {
	fmt.Println("Masukkan pilihan anda:")
	fmt.Println("1 - Add Account")
	fmt.Println("2 - Read Account")
	fmt.Println("3 - Update Account")
	fmt.Println("4 - Delete Account")
	fmt.Println("5 - Top-Up")
	fmt.Println("6 - Transfer")
	fmt.Println("7 - History Top-Up")
	fmt.Println("8 - History Transfer")
	var pilihan string
	fmt.Scanln(&pilihan)

	switch pilihan {
	case "1":
		// add account
		// input - phone, name, 
	case "2":
		//
	}
}
