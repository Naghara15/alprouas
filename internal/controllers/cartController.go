package Controllers

import (
	"alprouas/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cartResponse struct {
	Id		int		`json:"id"`
	Product_id	int	`json:"product_id"`
	Image 	string	`json:"image"`
	Name 	string	`json:"name"`
	Stock 	int		`json:"stock"`
	Qty		int		`json:"qty"`
	Price 	float64	`json:"price"`
}

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

	response := make([]cartResponse, len(carts))
	for i, cart := range carts {
		product, err := models.GetProductByID(cart.Product_id)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
			return
		}

		response[i].Id = cart.Id
		response[i].Product_id = cart.Product_id
		response[i].Image = product.Image
		response[i].Name = product.Name
		response[i].Stock = product.Stock
		response[i].Qty = cart.Qty
		response[i].Price = cart.Total_price
	}

	c.IndentedJSON(http.StatusOK, response)
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

func CalculateTotalHandler(c *gin.Context) {
	var carts []models.Cart
	err := c.BindJSON(&carts)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	total, err := models.CalculateTotal(carts)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.IndentedJSON(http.StatusOK, total)
}

func NotificationCartHandler(c *gin.Context) {
	var cart models.Cart
	err := c.BindJSON(&cart)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	carts, err := models.GetCartUser(cart.User_id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, len(carts))
}