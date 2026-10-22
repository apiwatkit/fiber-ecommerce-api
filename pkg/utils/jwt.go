package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims  struct {
	UserID uint `json:"user_id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint, role string) (string, error)  {
	// 1. อ่าน JWT Secret จาก environment
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET is not configured")
	}

	// 2. กำหนด Claims ที่ต้องการเก็บใน Token
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	// 3. สร้าง Token ด้วยอัลกอริทึม HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 4. Sign Token ด้วย Secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	// 5. ส่ง Token กลับไปให้ผู้เรียกใช้งาน
	return tokenString, nil
}

func ValidateJWT(tokenString string) (*JWTClaims, error) {
	// อ่าน Secret Key จาก environment
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return nil, errors.New("JWT_SECRET is not configured")
	}

	// Parse และตรวจสอบ Token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// ป้องกัน Token ที่ใช้ Algorithm อื่น
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(secretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	// ตรวจสอบว่า Token ผ่าน Validation หรือไม่
	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}