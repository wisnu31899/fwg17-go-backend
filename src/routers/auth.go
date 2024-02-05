package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/wisnu31899/fwg17-go-backend/src/controllers/auth"
)

func AuthRouter(rg *gin.RouterGroup) {
	rg.POST("/login", controllers.Login)
	rg.POST("/register", controllers.Register)
	rg.POST("/forgot-password", controllers.ForgotPassword)
}
