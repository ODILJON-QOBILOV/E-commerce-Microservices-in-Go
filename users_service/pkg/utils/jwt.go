package utils

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var secret = []byte("supersecretkey") // меняй на свой секрет в продакшене

// GenerateToken создаёт JWT токен с userID и временем жизни
func GenerateToken(userID int64, duration time.Duration) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(duration).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secret)
}

// ParseToken проверяет токен и возвращает userID
func ParseToken(tokenStr string) (int64, error) {
    token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return secret, nil
    })
    if err != nil {
        return 0, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        userID, ok := claims["user_id"].(float64)
        if !ok {
            return 0, errors.New("invalid token claims")
        }
        return int64(userID), nil
    }

    return 0, errors.New("invalid token")
}
