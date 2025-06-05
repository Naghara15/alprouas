package Controllers

import (
	"alprouas/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTagsHandler(c *gin.Context) {
	tags, err := models.GetAllTags()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, tags)
}

func AddClickHandler(c *gin.Context) {
	input := struct {
		User_id    int `json:"user_id"`
		Product_id int `json:"product_id"`
	}{}

	err := c.BindJSON(&input)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	// Buat slice User_tags yang mempunyai product id yang sama
	product_tags, err := models.GetTagsByProductID(input.Product_id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	// Membuat []User_tags sesuai dengan tag yang didapat diatas
	user_tags := make([]models.User_tag, len(product_tags))
	for i, s := range product_tags {
		user_tags[i].User_id = input.User_id
		user_tags[i].Tag_id = s.Tag_id
	}

	// Menambahkan user_tags.Clicked
	err = models.AddClick(user_tags)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, "OK")
}
