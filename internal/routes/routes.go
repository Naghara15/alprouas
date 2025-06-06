package routes

import (
	Controllers "alprouas/internal/controllers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	// User
	r.POST("/user", Controllers.CreateUserHandler)
	r.GET("/user", Controllers.GetAllUserHandler)

	// Login
	r.POST("/login", Controllers.LoginHandler)

	// Product
	r.POST("/products", Controllers.GetProductsByTagHandler)
	r.GET("/products", Controllers.GetAllProductsHandler)
	r.POST("/recommend", Controllers.GetRecommendation)
	

	// Tag
	r.GET("/tags", Controllers.GetAllTagsHandler)
	r.POST("/addclick", Controllers.AddClickHandler)

	// Cart
	r.POST("/getcart", Controllers.GetCartUserHandler)
	r.POST("/addcart", Controllers.AddCartHandler)
	r.DELETE("/deletecart", Controllers.DeleteCartHandler)
	r.POST("/updateqtycart", Controllers.UpdateCartQtyHandler)

	// Transaction
	r.POST("/buy", Controllers.BuyHandler)
}
