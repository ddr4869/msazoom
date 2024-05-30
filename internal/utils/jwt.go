package utils

import (
	"time"

	"github.com/ddr4869/msazoom/config"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/jwt"
)

// SdtClaims defines the custom claims
type SdtClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt_lib.StandardClaims
}

func GenerateToken(username string, role int) {
	// generate jwt token
	jwt.Auth("1234567890987654321")

}

// GenerateJWT generates token from the given information
func GenerateJWT(name, role string) (string, error) {
	claims := SdtClaims{
		name,
		role,
		jwt_lib.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    config.Issuer,
		},
	}

	token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JwtSecretPassword))

	return tokenString, err
}
