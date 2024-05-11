package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"maxprofit/internal/jwt"
)

const (
	authorizationHeaderKey = "x-authorization"
)

func JWTValidatorHandlerFunc(jwtValidator jwt.Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.GetHeader(authorizationHeaderKey)
		isValid, err := jwtValidator.ValidateToken(jwtToken)

		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("error occured while validating the token %w\n", err))
		}

		if !isValid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
