package calculatemaxprofit

import "maxprofit/internal/interactor/calculatemaxprofit"

type Controller struct {
	calculateMaxProfitInteractor *calculatemaxprofit.Interactor
}

func NewController(calculateMaxProfitInteractor *calculatemaxprofit.Interactor) *Controller {
	return &Controller{
		calculateMaxProfitInteractor: calculateMaxProfitInteractor,
	}
}
