package jwt

type Validator interface {
	ValidateToken(token string) (bool, error)
}
