package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/middlewares"
)

func Combine(r *gin.RouterGroup) {
	authMiddleware, _ := middlewares.Auth()
	r.Use(authMiddleware.MiddlewareFunc())
	UserRouter(r.Group("/user"))
	ProductRouter(r.Group("/products"))
	OrderRouter(r.Group("/order"))
}
