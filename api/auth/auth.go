package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Authenticated struct
type Authenticated struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

// HashPassword hashing password
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword compare hashed password with password string
func VerifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateToken generate JWT token
func GenerateToken(id uuid.UUID, additionalClaims jwt.MapClaims) (string, string, error) {
	apiSecret := []byte(os.Getenv("API_SECRET"))

	// AccessToken
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    id.String(),
		"exp":        time.Now().Add(time.Minute * 15).Unix(),
	}
	if len(additionalClaims) != 0 {
		for k, v := range additionalClaims {
			claims[k] = v
		}
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString(apiSecret)
	if err != nil {
		return "", "", err
	}

	// RefreshToken
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["user_id"] = id.String()
	refreshTokenClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	refresh, err := refreshToken.SignedString(apiSecret)
	if err != nil {
		return "", "", err
	}

	return token, refresh, nil
}

// JWTTokenValidate validate jwt token
func JWTTokenValidate(c echo.Context) error {
	tkn, err := ExtractToken(c)
	if err != nil {
		return err
	}
	token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

// ExtractToken extract token from Bearer
func ExtractToken(c echo.Context) (string, error) {
	bearerToken := c.Request().Header.Get("Authorization")
	splitted := strings.Split(bearerToken, " ")
	if len(splitted) == 2 {
		return splitted[1], nil
	}
	return "nil", fmt.Errorf("Bearer doesn't valid")
}

// Pretty string prettier
func Pretty(data interface{}) {
	_, err := json.MarshalIndent(data, "", "")
	if err != nil {
		return
	}
}
