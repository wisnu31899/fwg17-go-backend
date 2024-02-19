package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/middlewares"
)

func Combine(r *gin.RouterGroup) {
	authMiddleware, _ := middlewares.Auth()
	r.Use(authMiddleware.MiddlewareFunc())
	UserRouter(r.Group("/users"))
	ProductRouter(r.Group("/products"))
	ProductSizeRouter(r.Group("/productSize"))
	ProductVariantRouter(r.Group("/productVariant"))
	TagsRouter(r.Group("/tags"))
	PromoRouter(r.Group("/promo"))
	ProductTagRouter(r.Group("/productTags"))
	ProductRatingsRouter(r.Group("/productRatings"))
	ProductCategoriesRouter(r.Group("/productCategories"))
	OrdersRouter(r.Group("/orders"))
	OrderDetailsRouter(r.Group("/orderDetails"))
	CategoriesRouter(r.Group("/categories"))
	MessageRouter(r.Group("/message"))
}
