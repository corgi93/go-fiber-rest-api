package utils

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type TokenMetadata struct {
	Expires int64
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)

	if err != nil {
		return nil, err
	}
	return token, nil

}

// jwt에서 meta data 추출
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)

	if err != nil {
		return nil, err
	}

	// token, credentials 체킹하고 세팅
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// epxires time
		expires := int64(claims["exp"].(float64))

		return &TokenMetadata{
			Expires: expires,
		}, nil
	}
	return nil, err
}
