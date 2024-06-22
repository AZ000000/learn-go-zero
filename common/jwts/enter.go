package jwts

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtPayload struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
}

type CustomClaims struct {
	JwtPayload
	jwt.RegisteredClaims
}

func GenToken(user JwtPayload, accessSecret string, expires int64) (string, error) {
	clain := CustomClaims{
		JwtPayload: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expires))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clain)
	return token.SignedString([]byte(accessSecret))

}

func ParseToken(tokenString string, accessSecret string, expires int64) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if calims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return calims, err
	}
	return nil, errors.New("invalid token")
}
