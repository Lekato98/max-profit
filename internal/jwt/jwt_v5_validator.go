package jwt

import "github.com/golang-jwt/jwt/v5"

type V5 struct {
	secretKey string
}

func NewV5Validator(secretKey string) *V5 {
	return &V5{
		secretKey: secretKey,
	}
}

func (v5 *V5) ValidateToken(token string) (bool, error) {
	decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(v5.secretKey), nil
	})

	if err != nil {
		return false, err
	}

	return decodedToken.Valid, nil
}
