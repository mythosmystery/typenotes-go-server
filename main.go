package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mythosmystery/typenotes-go-server/config"
	"github.com/mythosmystery/typenotes-go-server/models"
	"github.com/mythosmystery/typenotes-go-server/routes"
	"github.com/mythosmystery/typenotes-go-server/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:Pepper113@tcp(127.0.0.1:3306)/typenotes?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	config.DB.AutoMigrate(&models.User{}, &models.Note{})
	fmt.Println("Connected to database")
}

func main() {
	app := fiber.New()

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	services.InitializeGmail()

	ConnectDB()
	routes.InitRoutes(app)

	app.Listen(":3000")
}
