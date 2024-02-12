package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers/admin"
)

func ProductRouter(r *gin.RouterGroup) {
	r.GET("", admin.GetAllProducts)
	r.GET("/:id", admin.GetDetailProduct)
	r.POST("", admin.CreateProduct)
	r.PATCH("/:id", admin.UpdateProduct)
	r.DELETE("/:id", admin.DeleteProduct)
}
