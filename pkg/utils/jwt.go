package utils

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var jwtKey []byte

func init() {
	// Load environment variables from the specified file
	if err := godotenv.Load("jwt-token.env"); err != nil {
		panic("Error loading jwt-token.env file")
	}

	// Load JWT key from environment variable
	jwtKey = []byte(os.Getenv("JWT_SECRET"))
	if len(jwtKey) == 0 {
		panic("JWT_SECRET environment variable is not set")
	}
}

// Claims представляет данные, которые будут храниться в токене
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// Генерация JWT токена
func GenerateJWT(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Проверка и парсинг токена
func ParseJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

// Проверка токена и извлечение данных
func VerifyToken(ctx context.Context) (string, error) {
	// Извлечение токена из метаданных (например, Authorization)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "no metadata found")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "no authorization token provided")
	}

	tokenStr := authHeader[0]
	claims, err := ParseJWT(tokenStr)
	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	return claims.UserID, nil
}
