package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers/admin"
)

func TagsRouter(r *gin.RouterGroup) {
	r.GET("", admin.GetAllTags)
	r.GET("/:id", admin.GetDetailTag)
	r.POST("", admin.CreateTag)
	r.PATCH("/:id", admin.UpdateTag)
	r.DELETE("/:id", admin.DeleteTag)
}
