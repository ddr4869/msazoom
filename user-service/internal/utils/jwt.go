package utils

import (
	// "github.com/ddr4869/msazoom/config"
	// jwt_lib "github.com/dgrijalva/jwt-go"

	"log"
	"net/http"

	"github.com/ddr4869/msazoom/user-service/config"
	"github.com/ddr4869/msazoom/user-service/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(name, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"role": role,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.JwtSecretPassword))
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return tokenString, nil
}

func ParseJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")
		if len(bearToken) < 7 || bearToken[:7] != "Bearer " {
			dto.NewErrorResponse(c, http.StatusBadRequest, nil, "invalid token")
			return
		}
		bearToken = bearToken[7:]
		token, err := jwt.ParseWithClaims(bearToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return []byte(config.JwtSecretPassword), nil
		})
		if err != nil {
			dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to parse token")
		} else if claims, ok := token.Claims.(*UserClaims); ok {
			c.Set("claims", claims)
		} else {
			dto.NewErrorResponse(c, http.StatusInternalServerError, err, "unknown claims type, cannot proceed")
		}
		c.Next()
	}
}
