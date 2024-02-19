package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers/admin"
)

func OrderDetailsRouter(r *gin.RouterGroup) {
	r.GET("", admin.GetAllOrderDetails)
	r.GET("/:id", admin.GetDetailOrderDetail)
	r.POST("", admin.CreateOrderDetail)
	r.PATCH("/:id", admin.UpdateOrderDetail)
	r.DELETE("/:id", admin.DeleteOrderDetail)
}
