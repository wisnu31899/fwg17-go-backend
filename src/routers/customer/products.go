package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/controllers/customer"
)

func ProductRouter(r *gin.RouterGroup) {
	r.GET("", customer.GetAllProducts)
	r.GET("/:id", customer.GetDetailProduct)
}
