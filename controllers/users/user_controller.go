package users

import (
	"github.com/gin-gonic/gin"
	"github.com/julsan26/bookstore-user-api/domain/users"
	"github.com/julsan26/bookstore-user-api/services"
	"github.com/julsan26/bookstore-user-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad request",
		}
		c.JSON(restErr.Status, restErr)
		return
	}
	result, savErr := services.CreateUser(user)

	if savErr != nil {
		c.JSON(savErr.Status, savErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Need to implement")
}
