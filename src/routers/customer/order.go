package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers/customer"
)

func OrderRouter(r *gin.RouterGroup) {
	r.POST("", customer.CreateOrder)
	r.GET("/:id", customer.GetDetailOrder)
}
