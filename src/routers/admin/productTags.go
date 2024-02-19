package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers/admin"
)

func ProductTagRouter(r *gin.RouterGroup) {
	r.GET("", admin.GetAllProductTag)
	r.GET("/:id", admin.GetDetailProductTag)
	r.POST("", admin.CreateProductTag)
	r.PATCH("/:id", admin.UpdateProductTag)
	r.DELETE("/:id", admin.DeleteProductTag)
}
