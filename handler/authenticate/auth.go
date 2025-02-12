package authenticate

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"iHR/repositories/model"
	"time"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	AuthType string `json:"auth_type"`
	Provider string `json:"provider"`
	jwt.RegisteredClaims
}

func NewAuth(secret string, authType string, provider string, userID uint, username string) (*model.Auth, error) {
	// For demo and dev purpose, we set it to a shorter time
	now := time.Now()
	tokenExpiredAt := now.Add(10 * time.Hour)
	token, err := GenerateToken(secret, authType, provider, tokenExpiredAt, now, userID, username)
	if err != nil {
		return nil, err
	}

	refreshTokenExpiredAt := now.Add(30 * 24 * time.Hour)
	refreshToken, err := GenerateToken(secret, authType, provider, refreshTokenExpiredAt, now, userID, username)
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

func GenerateToken(secret string, authType string, provider string, expiredAt time.Time, issuedAt time.Time, userID uint, username string) (string, error) {
	claims := &Claims{
		UserID:   userID,
		Username: username,
		AuthType: authType,
		Provider: provider,
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
