package Controllers

import (
	"alprouas/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var user models.Users

func CreateUserHandler(c *gin.Context) {
	err := c.BindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	if user.Username == "" || user.Password == "" {
		c.IndentedJSON(http.StatusBadRequest, "username or password cannot be empty")
		return
	}

	err = models.CreateUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	c.IndentedJSON(http.StatusOK, "user created")
}

func GetAllUserHandler(c *gin.Context) {
	users, err := models.GetAllUser()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return 
	}

	c.IndentedJSON(http.StatusOK, users)
}