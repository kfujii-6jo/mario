package utils

import (
	"time"
	"os"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(username string) (string, error) {
    secretKey := []byte(os.Getenv("SECRET_KEY"))
    expirationTime := time.Now().Add(time.Hour * 12)
    claims := &jwt.MapClaims{
        "username": username,
        "exp":      expirationTime.Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}
