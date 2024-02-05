package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/routers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	routers.Combine(r)
	r.Run()
}
