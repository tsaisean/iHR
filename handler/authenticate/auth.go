package authenticate

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"iHR/db/model"
	"time"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewAuth(secret string, userID uint, username string) (*model.Auth, error) {
	// For demo and dev purpose, we set it to a shorter time
	now := time.Now()
	tokenExpiredAt := now.Add(10 * time.Minute)
	token, err := GenerateToken(secret, userID, username, tokenExpiredAt, now)
	if err != nil {
		return nil, err
	}

	refreshTokenExpiredAt := now.Add(30 * 24 * time.Hour)
	refreshToken, err := GenerateToken(secret, userID, username, refreshTokenExpiredAt, now)
	if err != nil {
		return nil, err
	}

	auth := &model.Auth{
		UserID:                userID,
		Token:                 token,
		RefreshToken:          refreshToken,
		TokenExpiresAt:        tokenExpiredAt,
		RefreshTokenExpiresAt: refreshTokenExpiredAt,
		CreatedAt:             now,
	}

	return auth, nil
}

func GenerateToken(secret string, userID uint, username string, expiredAt time.Time, issuedAt time.Time) (string, error) {
	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredAt),
			IssuedAt:  jwt.NewNumericDate(issuedAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(secret string, tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
