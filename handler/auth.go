package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tshubham7/gorm-articles/repository"
	"github.com/tshubham7/gorm-articles/services"
)

// AuthHandler ..
type AuthHandler interface {
	// Create new account
	Register() gin.HandlerFunc

	// Login user
	Login() gin.HandlerFunc
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
				"message": "failed to create new account",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

// Login ...
func (au authHandler) Login() gin.HandlerFunc {
	sr := services.NewAuthService(au.u)
	return func(c *gin.Context) {
		// fetch credentials
		var params services.RegisterRequest
		err := c.Bind(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing or invalid params",
				"error":   err.Error(),
			})
			return
		}

		token, err := sr.Login(params)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "failed to login",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})

	}
}
