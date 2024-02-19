package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers/admin"
)

func PromoRouter(r *gin.RouterGroup) {
	r.GET("", admin.GetAllPromo)
	r.GET("/:id", admin.GetDetailPromo)
	r.POST("", admin.CreatePromo)
	r.PATCH("/:id", admin.UpdatePromo)
	r.DELETE("/:id", admin.DeletePromo)
}
