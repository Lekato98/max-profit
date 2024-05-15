package profit

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type profitInteractorSuit struct {
	suite.Suite
	interactor *Interactor
}

func (s *profitInteractorSuit) SetupTest() {
	s.interactor = NewInteractor()
}

func (s *profitInteractorSuit) TestNewInteractor() {
	s.T().Parallel()
	assert.NotNil(s.T(), s.interactor)
}

func TestSuit(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(profitInteractorSuit))
}
