package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers"
)

func TestimoniRouter(r *gin.RouterGroup) {
	r.GET("/join", controllers.ListAllTestimoniJoin)
	r.GET("/all", controllers.ListAllTestimoni)
	r.GET("/:id", controllers.DetailTestimoni)
	r.POST("", controllers.CreateTestimoni)
	r.PATCH("/:id", controllers.UpdateTestimoni)
	r.DELETE("/:id", controllers.DeleteTestimoni)
}
