package jwt

import "github.com/dgrijalva/jwt-go"

type SessionDetail struct {
	UserId    string `json:"user_id"`
	SessionId string `json:"session_id"`
}

type Claims struct {
	jwt.StandardClaims
	SessionDetail
}

