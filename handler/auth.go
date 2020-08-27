package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tshubham7/gorm-articles/repository"
	"github.com/tshubham7/gorm-articles/services"
)

// AuthHandler ..
type AuthHandler interface {
	Register() gin.HandlerFunc
}

type authHandler struct {
	u repository.UserService
}

//NewAuthHandler ..
func NewAuthHandler(u repository.UserService) AuthHandler {
	return authHandler{u}
}

// Register ...
func (au authHandler) Register() gin.HandlerFunc {
	sr := services.NewAuthService(au.u)

	return func(c *gin.Context) {
		params := services.RegisterRequest{}

		err := c.Bind(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "missing or invalid params",
			})
			return
		}

		token, err := sr.Register(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "something went wrong",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
