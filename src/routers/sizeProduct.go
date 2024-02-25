package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers"
)

func SizeProductGuestRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllPSByGuest)
}
