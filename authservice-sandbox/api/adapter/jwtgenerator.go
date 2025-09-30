package adapter

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("super-secret-key")

type JWTGenerator struct{}

func NewJWTGenerator() *JWTGenerator {
	return &JWTGenerator{}
}

func (j *JWTGenerator) GenerateAccessToken(username string) (string, error) {
	// TODO: Implementasi pembuatan token JWT
	// TODO: 1. Membuat Struktur Claims/Payload
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 15).Unix(), // Token berlaku selama 15 menit
	}

	//  TODO: 2. Signature JWT dengan algoritma HS256
	signature := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return signature.SignedString(secretKey)
}

func (j *JWTGenerator) ValidateToken(token string) (string, error) {
	// TODO: 1. Meneriman token dari user/client

	// TODO: 2. Memparsing token dengan secretKey
	tokenParsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	// TODO: 3. Mengektraksi username dari token yang sudah di parsing melalui claims
	if claims, ok := tokenParsed.Claims.(jwt.MapClaims); ok && tokenParsed.Valid {
		username := claims["username"].(string)
		return username, nil
	}
	return "", jwt.ErrSignatureInvalid
}
