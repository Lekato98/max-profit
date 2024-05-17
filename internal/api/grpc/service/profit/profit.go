package profit

import (
	calculatemaxprofitinteractor "maxprofit/internal/interactor/profit"
)

type ServerImpl struct {
	UnimplementedProfitServer
	calculateMaxProfitInteractor *calculatemaxprofitinteractor.Interactor
}

func NewProfitServer(calculateMaxProfitInteractor *calculatemaxprofitinteractor.Interactor) *ServerImpl {
	return &ServerImpl{
		calculateMaxProfitInteractor: calculateMaxProfitInteractor,
	}
}
