package Controllers

import (
	"alprouas/internal/models"
	"alprouas/internal/utilities"
	"net/http"

	"fmt"
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
		// Set cookie buat user id
		c.SetCookie("userid", fmt.Sprintf("%d", DataUser.Id), 3600, "/", "localhost", false, false)
		c.SetCookie("saldo", fmt.Sprintf("%.2f", DataUser.Saldo), 3600, "/", "localhost", false, false)
		c.SetCookie("name", DataUser.Username, 3600, "/", "localhost", false, false)
		c.IndentedJSON(http.StatusOK, true)
	} else {
		c.IndentedJSON(http.StatusUnauthorized, false)
	}
}
