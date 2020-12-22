package config

import (
	"fmt"
	"golang_api/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SetupDatabaseConnection is a function to connect to database
func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("failed to load env file")
	}

	//ini gausah dipake kalo di lepi gede gabisa
	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	// dbHost := os.Getenv("DB_HOST")
	// dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprint("root:@tcp(localhost:3306)/golang_api_wgtik?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to make a connection with database")
	}

	//todo : simpan logic disini
	//kita simpan modelnya disini
	db.AutoMigrate(&entity.Book{}, &entity.User{})
	return db
}

// CloseDatabaseConnection is used to close database connection
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("failed to close connection with database")
	}
	dbSQL.Close()
}
