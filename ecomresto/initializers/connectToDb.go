package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnecttoDb() {
	var err error
	dsn := "postgresql://postgres:newpass@localhost/ecomresto?sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	} else {
		log.Println("Successfully connected to the database")
	}
}
