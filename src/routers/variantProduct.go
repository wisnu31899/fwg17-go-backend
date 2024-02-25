package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers"
)

func VariantProductGuestRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllPVByGuest)
}
