package jwt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	secretKey = "secret_key"
)

type jwtV5ValidatorSuit struct {
	suite.Suite
	validator *V5
}

func TestV5(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(jwtV5ValidatorSuit))
}

func (s *jwtV5ValidatorSuit) SetupTest() {
	s.validator = NewV5Validator(secretKey)
}

func (s *jwtV5ValidatorSuit) TestNewV5Validator() {
	s.T().Parallel()
	s.Assert().NotNil(NewV5Validator(secretKey))
}

func (s *jwtV5ValidatorSuit) TestV5_ValidateToken() {
	s.T().Parallel()
	// Arrange
	tests := []struct {
		name            string
		givenToken      string
		expectedErr     error
		expectedIsValid bool
	}{
		{
			name:            "invalid token malformed",
			givenToken:      "random_token",
			expectedErr:     fmt.Errorf("non empty err"),
			expectedIsValid: false,
		},
		{
			name:            "invalid token expired",
			givenToken:      "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTU4MTE0NDcsImp0aSI6InVzZXJfaWQifQ.LptD7nlrf1cU4mN9VgQbI8_oXHZm4vadIMIKbCfe_WiJZKWWvMFGUNNF1mQqGxa8ViqVOEdL6HtS3SzWaq3t_Q",
			expectedErr:     fmt.Errorf("non empty err"),
			expectedIsValid: false,
		},
		{
			name:            "valid token",
			givenToken:      "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMTU4OTc3MTAsImp0aSI6InVzZXJfaWQifQ.3YozrpXsTC-zL35wHdlBUCQd82VJXaCnNS_VurObJlPslwy_BPD0Wx4ejSg2dq0k_dDRK1cX2SG-P_laqrTO0A",
			expectedErr:     nil,
			expectedIsValid: true,
		},
	}

	for _, tc := range tests {
		s.T().Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// Act
			result, err := s.validator.ValidateToken(tc.givenToken)
			s.Assert().True((err == nil) == (tc.expectedErr == nil))
			s.Assert().Equal(tc.expectedIsValid, result)
		})
	}
}
