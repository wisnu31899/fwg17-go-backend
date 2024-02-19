package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers/admin"
)

func CategoriesRouter(r *gin.RouterGroup) {
	r.GET("", admin.GetAllCategories)
	r.GET("/:id", admin.GetDetailCategories)
	r.POST("", admin.CreateCategories)
	r.PATCH("/:id", admin.UpdateCategories)
	r.DELETE("/:id", admin.DeleteCategories)
}
