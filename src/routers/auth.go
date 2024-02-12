package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/wisnu31899/fwg17-go-backend/src/controllers/auth"
	"github.com/wisnu31899/fwg17-go-backend/src/middlewares"
)

func AuthRouter(rg *gin.RouterGroup) {
	authMiddleware, _ := middlewares.Auth()
	rg.POST("/login", authMiddleware.LoginHandler)
	rg.POST("/register", controllers.Register)
	rg.POST("/forgot-password", controllers.ForgotPassword)
}
