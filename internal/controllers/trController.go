package Controllers

import (
	"alprouas/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuyHandler(c *gin.Context) {
	var cartIds []models.Cart
	err := c.BindJSON(&cartIds)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	carts, err := models.GetCartsByID(cartIds)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	for _,s := range carts{
		avail,err := CheckStock(s.Product_id)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError,err)
			return
		}
		if !avail {
			c.IndentedJSON(http.StatusConflict, "stock habis")
			return
		}
	}

	err = models.Buy(carts)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}


	c.IndentedJSON(http.StatusOK, "transaction complete")
}