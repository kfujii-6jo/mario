package utils

import (
	"time"
	"os"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// var secretKey = []byte("673c1a725c548b48fe1568045125b1d3c88ecd81f9f68994ce4ff09d2f3dc261")

func GenerateJWT(username string) (string, error) {
	err := godotenv.Load()
    if err != nil {
	    return "", err
    }
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
