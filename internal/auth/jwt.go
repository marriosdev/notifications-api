package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

type CustomClaims struct {
	UserID   string   `json:"uid"`
	Username string   `json:"usr"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

type UserAuthPayload struct {
	ID       string
	Username string
}

func GenerateToken(u UserAuthPayload) (string, error) {
	now := time.Now()
	claims := CustomClaims{
		UserID:   u.ID,
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "marrios.api",
			Subject:   u.Username,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(15 * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	tok, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("algoritmo inesperado")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tok.Claims.(*CustomClaims)
	if !ok || !tok.Valid {
		return nil, errors.New("token inválido")
	}
	return claims, nil
}
