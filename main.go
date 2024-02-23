package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wisnu31899/fwg17-go-backend/src/routers"
	"github.com/wisnu31899/fwg17-go-backend/src/services"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4000"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type, Authorization"},
	}))
	r.Static("/uploads", "./uploads")

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	routers.Combine(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &services.ResponseOnly{
			Success: false,
			Message: "resouce not found and retry again",
		})
	})
	r.Run(":5050")
}
