package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers/customer"
)

func UserRouter(r *gin.RouterGroup) {
	r.GET("/", customer.GetDetailUser)
	r.PATCH("/", customer.UpdateUser)
}
