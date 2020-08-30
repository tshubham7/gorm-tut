package middleware

import (
	"log"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"github.com/tshubham7/gorm-articles/services"
)

//Verifier verify token
func Verifier() func(http.Handler) http.Handler {
	return jwtauth.Verifier(services.Token())
}

//Authenticate is used to validate the authentication of an user
//if the token is missing an error will accurd
func Authenticate(r *http.Request) error {
	token, _, err := jwtauth.FromContext(r.Context())
	if err != nil || token == nil {
		return errors.Wrap(err, "Not Authorized")
	}
	if !token.Valid {
		return errors.New("Invalid or expired token")
	}
	return nil
}

// Authenticator is the middleware that checks
// if the request containts an auth token
func Authenticator() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := Authenticate(r)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				render.JSON(w, r, map[string]string{"error": err.Error()})
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// UserID ... get user id from claim
func UserID(r *http.Request) string {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		log.Printf("Error getting id from claims, err:%v", err)
	}
	return claims["id"].(string)
}
