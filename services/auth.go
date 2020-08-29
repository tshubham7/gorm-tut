package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/tshubham7/gorm-articles/models"
	"github.com/tshubham7/gorm-articles/repository"
	"golang.org/x/crypto/bcrypt"
)

// Authtoken ..
type Authtoken struct {
	Token        string    `json:"token"`
	Expires      time.Time `json:"expires"`
	RefreshToken string    `json:"refreshToken"`
}

// RegisterRequest ...
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type auth struct {
	u repository.UserService
}

// AuthService ...
type AuthService interface {
	// create new user
	Register(user RegisterRequest) (Authtoken, error)

	// login user
	Login(user RegisterRequest) (Authtoken, error)

	// check if password is correct
	IsPasswordValid(hash string, password string) bool
}

// NewAuthService ...
func NewAuthService(u repository.UserService) AuthService {
	jawt = jwtauth.New("HS256", []byte(getSecret()), nil)
	return &auth{u}
}

func getSecret() string {
	// use environment variable here
	secret := "L2A0O4D7A8L4E1L6E3"
	if secret == "" {
		panic("SECRET env Missing")
	}

	return secret
}

var jawt *jwtauth.JWTAuth

func token() *jwtauth.JWTAuth {
	return jawt
}

// sign sings the token with id and email for later use
func sign(ID string, email string) (Authtoken, error) {
	expires := time.Now().Add(time.Hour * time.Duration(24*7)) // set to 10 min
	claims := jwt.MapClaims{}
	jwtauth.SetIssuedAt(claims, time.Now())
	jwtauth.SetExpiryIn(claims, time.Hour*time.Duration(24*7)) // set to 10 min
	claims["id"] = ID
	claims["role"] = "auth"
	if email != "" {
		claims["email"] = email
	}

	_, tokenString, err := token().Encode(claims)

	claims["role"] = "refresh"
	jwtauth.SetExpiryIn(claims, time.Hour*time.Duration(24*7))
	_, refreshTokenString, err := token().Encode(claims)
	return Authtoken{
		Token:        tokenString,
		Expires:      expires,
		RefreshToken: refreshTokenString,
	}, err
}

// IsPasswordValid ...
func (a auth) IsPasswordValid(hash string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err == nil {
		return true
	}
	return false
}

// Register ...
func (a auth) Register(user RegisterRequest) (Authtoken, error) {
	um := user.ToModel()
	err := a.u.Create(um).Scan(um).Error
	if err != nil {
		return Authtoken{}, err
	}
	return sign(um.ID.String(), um.Email)
}

// Login ...
func (a auth) Login(user RegisterRequest) (Authtoken, error) {
	um, err := a.u.GetByEmail(user.Email)
	if err != nil {
		return Authtoken{}, fmt.Errorf("Not user found with email %s", user.Email)
	}

	// checking password
	if !a.IsPasswordValid(um.Password, user.Password) {
		return Authtoken{}, errors.New("Invalid password")
	}

	return sign(um.ID.String(), um.Email)

}

// HashPassword generate a hash from password
// also use to hash the token for password reset
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// ToModel ...
func (ur RegisterRequest) ToModel() *models.User {
	hashed, _ := hashPassword(ur.Password)
	return &models.User{
		Name:     ur.Name,
		Email:    ur.Email,
		Password: hashed,
	}
}
