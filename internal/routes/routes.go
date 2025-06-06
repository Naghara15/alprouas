package routes

import (
	Controllers "alprouas/internal/controllers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	// User
	r.POST("/user", Controllers.CreateUserHandler)
	r.GET("/user", Controllers.GetAllUserHandler)
	r.PUT("/updatesaldo", Controllers.UpdateSaldoHandler)

	// Login
	r.POST("/login", Controllers.LoginHandler)

	// Product
	r.POST("/products", Controllers.GetProductsByTagHandler)
	r.GET("/products", Controllers.GetAllProductsHandler)
	r.POST("/recommend", Controllers.GetRecommendation)
	r.POST("/searchproduct", Controllers.SearchProductsHandler)

	// Tag
	r.GET("/tags", Controllers.GetAllTagsHandler)
	r.POST("/addclick", Controllers.AddClickHandler)

	// Cart
	r.POST("/getcart", Controllers.GetCartUserHandler)
	r.POST("/addcart", Controllers.AddCartHandler)
	r.DELETE("/deletecart", Controllers.DeleteCartHandler)
	r.PUT("/updateqtycart", Controllers.UpdateCartQtyHandler)
	r.POST("/calculatetotal", Controllers.CalculateTotalHandler)

	// Transaction
	r.POST("/buy", Controllers.BuyHandler)
}
