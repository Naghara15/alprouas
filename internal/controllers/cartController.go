package Controllers

import (
	"alprouas/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCartUserHandler(c *gin.Context) {
	var cart models.Cart

	err := c.BindJSON(&cart)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	carts, err := models.GetCartUser(cart.User_id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, carts)
}

func AddCartHandler(c *gin.Context) {
	var cart models.Cart
	var err error
	err = c.BindJSON(&cart)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	product, err := models.GetProductByID(cart.Product_id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	if product.Stock <= 0 {
		c.IndentedJSON(http.StatusConflict, "Stock habis")
		return
	}

	cart.Total_price = product.Price * float64(cart.Qty)

	err = models.AddCart(cart)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	carts, err := models.GetCartUser(cart.User_id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, carts)
}

func DeleteCartHandler(c *gin.Context) {
	var cart models.Cart
	var err error

	err = c.BindJSON(&cart)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	err = models.DeleteCart(cart.Id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	carts, err := models.GetCartUser(cart.User_id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, carts)
}

func UpdateCartQtyHandler(c *gin.Context) {
	var cart models.Cart
	err := c.BindJSON(&cart)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = models.UpdateCartQty(cart)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, "OK")
}