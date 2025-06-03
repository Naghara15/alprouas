package Controllers

import (
	"alprouas/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetProductsByTagHandler(c *gin.Context) {
	var relation []models.Product_tag

	err := c.BindJSON(&relation)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	tags := make([]int, len(relation))
	for i, s := range relation {
		tags[i] = s.Tag_id
	}

	products, err := models.GetProductsByTag(tags)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, products)
}

func GetAllProductsHandler(c *gin.Context) {
	products, err := models.GetAllProducts()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, products)
}

func GetRecommendation(c *gin.Context) {
	var user models.User_tag

	err := c.BindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	user_tags, err := models.GetUserMostClicked(user.User_id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	tags := make([]int, len(user_tags))
	for i, s := range user_tags {
		tags[i] = s.Tag_id
		fmt.Println(tags[i])
	}

	products, err := models.GetProductsByTag(tags)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, products)
}

func CheckStock(id int) (bool, error) {
	product, err := models.GetProductByID(id)
	if err != nil {
		return false, err
	}

	if product.Stock <= 0 {
		return false, nil
	}
	return true, nil
}