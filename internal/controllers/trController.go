package Controllers

import (
	"alprouas/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuyHandler(c *gin.Context) {
	var cartIds []models.Cart
	err := c.BindJSON(&cartIds)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	carts, err := models.GetCartsByID(cartIds)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var total float64
	for _, s := range carts {
		avail, err := CheckStock(s.Product_id)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
			return
		}
		if !avail {
			c.IndentedJSON(http.StatusConflict, "stock habis")
			return
		}
		total += s.Total_price
	}

	user, err := models.GetuserByID(carts[0].User_id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if user.Saldo < total {
		c.IndentedJSON(http.StatusPaymentRequired, "saldo tidak cukup")
		return
	}

	err = models.Buy(carts)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	user.Saldo -= total
	err = models.UpdateSaldo(user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.SetCookie("totalbayar", fmt.Sprintf("%.2f", total), 3600, "/", "localhost", false, false)
	c.SetCookie("saldo", fmt.Sprintf("%.2f", user.Saldo), 3600, "/", "localhost", false, false)

	c.IndentedJSON(http.StatusOK, "transaction complete")
}
