package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/routers"
	"github.com/wisnu31899/fwg17-go-backend/src/services"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4040"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	}))
	routers.Combine(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &services.ResponseOnly{
			Success: false,
			Message: "resouce not found and retry again",
		})
	})
	r.Run("127.0.0.1:5050")
}
