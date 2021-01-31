package schema

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	Issued   int64
	Identity string
}

var jwtSecret = os.Getenv("JWT_SECRET")

func CreateJWT(identity string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"issued":   time.Now().Unix(),
		"identity": identity,
	})
	return token.SignedString("SentimentAPISecret")
}

func ReadJWT(tokenString string) (JwtClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		return JwtClaims{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return JwtClaims{Issued: claims["issued"].(int64), Identity: claims["identity"].(string)}, nil
	}
	return JwtClaims{}, fmt.Errorf("Token is invalid")
}
