package supabase

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func Connect() *gorm.DB {
	host := os.Getenv("SUPABASE_DB_HOST")
	user := os.Getenv("SUPABASE_DB_USER")
	password := os.Getenv("SUPABASE_DB_PASSWORD")
	dbName := os.Getenv("SUPABASE_DB_NAME")
	port := os.Getenv("SUPABASE_DB_PORT")
	timeZone := os.Getenv("TIMEZONE")
	sslmode := "disable"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbName, port, sslmode, timeZone)
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		log.Fatal(e)
	}

	return database
}

func DB() *gorm.DB {
	return database
}
