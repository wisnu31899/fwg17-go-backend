package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers/admin"
)

func MessageRouter(r *gin.RouterGroup) {
	r.GET("", admin.GetAllMessages)
	r.GET("/:id", admin.GetDetailMessage)
	r.POST("", admin.CreateMessage)
	r.PATCH("/:id", admin.UpdateMessage)
	r.DELETE("/:id", admin.DeleteMessage)
}
