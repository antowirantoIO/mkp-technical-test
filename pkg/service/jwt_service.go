package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"mkp-boarding-test/internal/domain/entity"
)

type JWTService interface {
	GenerateToken(user *entity.User) (string, error)
	GenerateRefreshToken(user *entity.User) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	ExtractUserIDFromToken(tokenString string) (string, error)
	ExtractClaimsFromToken(tokenString string) (jwt.MapClaims, error)
}

type jwtService struct {
	secretKey    string
	refreshKey   string
	tokenExpiry  time.Duration
	refreshExpiry time.Duration
}

type JWTClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	TokenID  string `json:"token_id"`
	jwt.RegisteredClaims
}

func NewJWTService(secretKey, refreshKey string, tokenExpiry, refreshExpiry time.Duration) JWTService {
	return &jwtService{
		secretKey:     secretKey,
		refreshKey:    refreshKey,
		tokenExpiry:   tokenExpiry,
		refreshExpiry: refreshExpiry,
	}
}

func (j *jwtService) GenerateToken(user *entity.User) (string, error) {
	claims := JWTClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		TokenID:  uuid.New().String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.tokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jwtService) GenerateRefreshToken(user *entity.User) (string, error) {
	claims := JWTClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		TokenID:  uuid.New().String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.refreshExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.refreshKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) ExtractUserIDFromToken(tokenString string) (string, error) {
	token, err := j.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, exists := claims["user_id"]; exists {
			if userIDStr, ok := userID.(string); ok {
				return userIDStr, nil
			}
		}
		return "", errors.New("user_id not found in token")
	}

	return "", errors.New("invalid token")
}

func (j *jwtService) ExtractClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := j.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}