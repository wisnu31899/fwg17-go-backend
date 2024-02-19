package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers/admin"
)

func OrdersRouter(r *gin.RouterGroup) {
	r.GET("", admin.GetAllOrders)
	r.GET("/:id", admin.GetDetailOrder)
	r.POST("", admin.CreateOrder)
	r.PATCH("/:id", admin.UpdateOrder)
	r.DELETE("/:id", admin.DeleteOrder)
}
