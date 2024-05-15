package profit

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (s *profitInteractorSuit) TestInteractor_CalculateMaxProfit() {
	s.T().Parallel()
	// Arrange
	tests := []struct {
		name           string
		givenProfits   []int64
		expectedErr    error
		expectedResult int64
	}{
		{
			name:           "passing empty profits should return 0",
			givenProfits:   []int64{},
			expectedErr:    nil,
			expectedResult: 0,
		},
		{
			name:           "passing non empty profits should return first element",
			givenProfits:   []int64{10, 20, 30},
			expectedErr:    nil,
			expectedResult: 10,
		},
	}

	for _, tc := range tests {
		s.T().Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// Act
			result, err := s.interactor.CalculateMaxProfit(context.TODO(), tc.givenProfits)

			// Assert
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
