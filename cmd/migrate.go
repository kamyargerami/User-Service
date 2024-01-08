package main

import (
	"fmt"
	"github.com/kam2yar/user-service/internal/database/connections"
	"github.com/kam2yar/user-service/internal/database/entities"
	"time"
)

func main() {
	fmt.Println("Start migrating database structures", time.RFC822)
	migrate()
	fmt.Println("Migrations finished successfully", time.RFC822)
}

func migrate() {
	db := connections.DefaultConnection()

	fmt.Print("Processing User... ")
	db.AutoMigrate(&entities.User{})
	fmt.Println("(Done)")
}