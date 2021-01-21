package main

import (
	"os"
	"sharavara/test-api/controllers"
	"sharavara/test-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	dbHost, dbPort, dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	models.ConnectDatabase(dbHost, dbPort, dbUser, dbPassword, dbName)

	// Routes
	r.GET("/", controllers.Health)
	r.GET("/ping", controllers.Ping)
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// Run the server
	r.Run(":8080")
}
