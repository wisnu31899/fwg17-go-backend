package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/routers/admin"
	"github.com/wisnu31899/fwg17-go-backend/src/routers/customer"
)

func Combine(r *gin.Engine) {
	AuthRouter(r.Group("/auth"))
	admin.Combine(r.Group("/admin"))
	customer.Combine(r.Group("/customer"))
	ProductGuestRouter(r.Group("/products"))
	TestimoniRouter(r.Group("/testimoni"))
	SizeProductGuestRouter(r.Group("/sizeproducts"))
	VariantProductGuestRouter(r.Group("/variantproducts"))
}
