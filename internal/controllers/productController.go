package Controllers

import (
	"alprouas/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SearchProductsHandler(c *gin.Context) {
	var product models.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	products, err := models.SearchProducts(product.Name)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, products)
}

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

func GetProductIDHandler(c *gin.Context) {
	productId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	product, err := models.GetProductByID(productId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, product)
}