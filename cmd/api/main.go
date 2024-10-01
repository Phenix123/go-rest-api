package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	_ "testApi/cmd/api/docs" // docs are generated by Swag CLI, you have to import it.
	"testApi/internal/api"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found")
	}
}

//	@title			Example REST API server for albums
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Developer
//	@contact.email	d.kurlayk@monodigital.ru

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	gin.SetMode(gin.DebugMode)
	server := api.New("localhost", 8080)
	server.Init()
}
