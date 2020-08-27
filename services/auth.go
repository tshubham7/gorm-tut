package services

import (
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
	Register(user RegisterRequest) (Authtoken, error)
}

// NewAuthService ...
func NewAuthService(u repository.UserService) AuthService {
	jawt = jwtauth.New("HS256", []byte(getSecret()), nil)
	return &auth{u}
}

func getSecret() string {
	// secret := os.Getenv("SECRET")
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

func (a *auth) Register(user RegisterRequest) (Authtoken, error) {
	um := user.ToModel()
	err := a.u.Create(um).Scan(um).Error
	if err != nil {
		return Authtoken{}, err
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
