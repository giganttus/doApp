package auth

import (
	"doApp/helpers"
	"doApp/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwt(userID int) (map[string]string, error) {
	accessKey := os.Getenv("ACCESS_KEY")
	if accessKey == "" {
		return nil, helpers.ErrEnvLoad
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, models.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{},
		UserID:           userID,
		ExpiresAt:        time.Now().AddDate(0, 6, 0).Unix(),
	}).SignedString([]byte(accessKey))

	if err != nil {
		return nil, helpers.ErrAuth
	}

	return map[string]string{
		"access_token": accessToken,
	}, nil

}

func ValidateJwt(tokenParam string, tokenKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenParam, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})

	if err != nil {
		return nil, helpers.ErrAuth
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, helpers.ErrAuth
	}
}
