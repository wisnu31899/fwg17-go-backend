package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers/admin"
)

func ProductRatingsRouter(r *gin.RouterGroup) {
	r.GET("", admin.GetAllProductRatings)
	r.GET("/:id", admin.GetDetailProductRating)
	r.POST("", admin.CreateProductRating)
	r.PATCH("/:id", admin.UpdateProductRating)
	r.DELETE("/:id", admin.DeleteProductRating)
}
