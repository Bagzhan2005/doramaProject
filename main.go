package main

import (
	"doramaPro/database"
	"doramaPro/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	dramaHandler := handlers.NewDramaHandler(db)

	api := router.Group("/api")
	{
		api.GET("/dramas", dramaHandler.GetDramas)
		api.POST("/dramas", dramaHandler.CreateDrama)
		api.PUT("/dramas/:id", dramaHandler.UpdateDrama)
		api.DELETE("/dramas/:id", dramaHandler.DeleteDrama)
	}

	router.Run(":8080")
}
