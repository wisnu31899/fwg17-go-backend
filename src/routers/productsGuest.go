package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers"
)

func ProductGuestRouter(r *gin.RouterGroup) {
	r.GET("", controllers.GetAllProducts)
	r.GET("/:id", controllers.GetDetailProduct)
}
