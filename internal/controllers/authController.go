package Controllers

import (
	"alprouas/internal/models"
	"alprouas/internal/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var user models.Users

	err := c.BindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	DataUser, err := models.GetUserByName(user.Username)
	if err != nil && DataUser.Id == 0 {
		c.IndentedJSON(http.StatusUnauthorized, err.Error())
		return
	}

	if utilities.VerifyHash(user.Password, DataUser.Password) {
		c.IndentedJSON(http.StatusOK, DataUser.Id)
	} else {
		c.IndentedJSON(http.StatusUnauthorized, false)
	}
}