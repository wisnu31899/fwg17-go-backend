package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers/admin"
)

func ProductSizeRouter(r *gin.RouterGroup) {
	r.GET("", admin.GetAllProductSize)
	r.GET("/:id", admin.GetDetailProductSize)
	r.POST("", admin.CreateProductSize)
	r.PATCH("/:id", admin.UpdateProductSize)
	r.DELETE("/:id", admin.DeleteProductSize)
}
