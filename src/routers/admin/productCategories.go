package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers/admin"
)

func ProductCategoriesRouter(r *gin.RouterGroup) {
	r.GET("", admin.GetAllProductCategories)
	r.GET("/:id", admin.GetDetailProductCategories)
	r.POST("", admin.CreateProductCategories)
	r.PATCH("/:id", admin.UpdateProductCategories)
	r.DELETE("/:id", admin.DeleteProductCategories)
}
