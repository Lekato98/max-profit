package profit

import "maxprofit/internal/interactor/profit"

type Controller struct {
	calculateMaxProfitInteractor *profit.Interactor
}

func NewController(calculateMaxProfitInteractor *profit.Interactor) *Controller {
	return &Controller{
		calculateMaxProfitInteractor: calculateMaxProfitInteractor,
	}
}
