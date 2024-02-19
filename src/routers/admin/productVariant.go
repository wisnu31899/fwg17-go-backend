package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers/admin"
)

func ProductVariantRouter(r *gin.RouterGroup) {
	r.GET("", admin.GetAllProductVariant)
	r.GET("/:id", admin.GetDetailProductVariant)
	r.POST("", admin.CreateProductVariant)
	r.PATCH("/:id", admin.UpdateProductVariant)
	r.DELETE("/:id", admin.DeleteProductVariant)
}
