package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

const (
	authorizationKey = "x-authorization"
)

func JWTValidatorFuncHandler(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.GetHeader(authorizationKey)
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (any, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.String(http.StatusUnauthorized, "unauthorized: invalid token")
			return
		}

		c.Next()
	}
}
