package Controllers

import (
	"net/http"
	"alprouas/internal/models"

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