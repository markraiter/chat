package handler

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	authHeader = "Authorization"
)

func (h *Handler) JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get(authHeader)
		if authHeader == "" {
			return c.String(echo.ErrUnauthorized.Code, "authorization header missing")
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return []byte("secret"), nil
		})

		if err != nil {
			return c.String(echo.ErrUnauthorized.Code, err.Error())
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := uint(claims["id"].(float64))
			c.Set("user_id", userID)

			return next(c)
		} else {
			return c.JSON(echo.ErrUnauthorized.Code, map[string]interface{}{
				"error": "invalid token",
			})
		}
	}
}
