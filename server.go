package main

import (
	_ "doApp/docs"
	"doApp/helpers"
	"doApp/routes"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title           Daily Offer API documentation
// @version         1.0
// @description     Rest API for mobile/web app

// @contact.name   API Support
// @contact.email  domagojpr@gmail.com

// @host      127.0.0.1:1312
// @BasePath  /api
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(helpers.ErrEnvLoad)
	}

	connString := os.Getenv("CON_STRING")
	if connString == "" {
		panic(helpers.ErrEnvLoad)
	}

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		panic(helpers.ErrDbConn)
	}

	routes := routes.InitRoutes(db)

	routes.Run("127.0.0.1:1312")
}
