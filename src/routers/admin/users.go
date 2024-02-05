package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers/admin"
)

func UserRouter(r *gin.RouterGroup) {
	r.GET("", admin.GetAllUsers)
	r.GET("/:id", admin.GetDetailUser)
	r.POST("", admin.CreateUser)
	r.PATCH("/:id", admin.UpdateUser)
	r.DELETE("/:id", admin.DeleteUser)
}
