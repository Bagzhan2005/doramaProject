package main

import (
	"doramaPro/database"
	"doramaPro/handlers"
	"doramaPro/middleware"
	"doramaPro/services"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	dramaService := services.NewDramaService(db)

	r := gin.Default()

	auth := handlers.NewAuthHandler(db)
	drama := handlers.NewDramaHandler(dramaService)

	api := r.Group("/api")
	{
		api.POST("/register", auth.Register)
		api.POST("/login", auth.Login)
		api.GET("/dramas", drama.GetDramas)

		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
		admin.POST("/dramas", drama.CreateDrama)
		admin.PUT("/dramas/:id", drama.UpdateDrama)
		admin.DELETE("/dramas/:id", drama.DeleteDrama)
	}

	r.Run(":8080")
}
