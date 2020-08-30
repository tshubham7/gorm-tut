package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/tshubham7/gorm-articles/services"
)

//Authenticate is used to validate the authentication of an user
//if the token is missing an error will accurd
func Authenticate(c *gin.Context) error {
	// r.Context().Value()
	tokenString, err := fetchHeaderAuthToken(c)

	if err != nil {
		return err
	}

	token, err := services.Validate(tokenString)

	if err != nil {
		return err
	}

	if token == nil || !token.Valid {
		return errors.New("Invalid or expired token")
	}
	return nil
}

// Authenticator is the middleware that checks
// if the request containts an auth token
func Authenticator() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := Authenticate(c)
		if err != nil {
			// stop right here
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
			return
		}
	}
}

// get header auth token from request
func fetchHeaderAuthToken(c *gin.Context) (string, error) {
	base := "Bearer "
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return "", errors.New("Invalid or expired token")
	}
	return tokenString[len(base):], nil
}

// UserID ... get user id from claim
func UserID(c *gin.Context) string {
	tokenString, err := fetchHeaderAuthToken(c)
	token, err := services.Validate(tokenString)

	if err != nil {
		log.Printf("Error getting id from claims, err:%v", err)
		return ""
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims["id"].(string)
}
