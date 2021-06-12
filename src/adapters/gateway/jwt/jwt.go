package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtInstance struct {
	Key string
}

func LoadJwt(key string) *JwtInstance {
	return &JwtInstance{
		Key: key,
	}
}

func (j *JwtInstance) ParseToken(token string) (*Claims, error) {
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Key), nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (j *JwtInstance) GenerateJWT(session SessionDetail) (string, error) {
	expirationTime := time.Now().Add(5 * 24 * time.Hour)

	claims := &Claims{
		SessionDetail: session,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.Key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

