package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Password yang ingin di-hash
	password := "admin123"

	// Generate hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("=== Password Hash Generator ===")
	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Hashed:   %s\n\n", string(hashedPassword))

	// SQL Insert query
	fmt.Println("=== SQL Insert Query ===")
	fmt.Printf(`INSERT INTO users (email, password, name, role) VALUES 
('admin@portfolio.com', '%s', 'Admin', 'admin');
`, string(hashedPassword))

	// Verify the hash works
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		fmt.Println("\n❌ Verification failed!")
	} else {
		fmt.Println("\n✅ Hash verification successful!")
	}
}
