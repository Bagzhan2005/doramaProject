package main

import (
	"doramaPro/database"
	"doramaPro/handlers"
	"doramaPro/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	auth := handlers.NewAuthHandler(db)
	drama := handlers.NewDramaHandler(db)

	api := r.Group("/api")
	{
		api.POST("/register", auth.Register)
		api.POST("/login", auth.Login)

		api.GET("/dramas", drama.GetDramas)

		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
		admin.POST("/dramas", drama.CreateDrama)
	}

	r.Run(":8080")
}
