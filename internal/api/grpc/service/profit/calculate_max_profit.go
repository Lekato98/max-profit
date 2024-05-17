package profit

import (
	"context"
)

// TODO protect the endpoint with JWT validator check where we need to pass it
func (c *ServerImpl) CalculateMaxProfit(ctx context.Context, profits *Profits) (*MaxProfit, error) {
	maxProfit, err := c.calculateMaxProfitInteractor.CalculateMaxProfit(ctx, profits.GetValues())
	if err != nil {
		return nil, err
	}

	return &MaxProfit{
		Value: maxProfit,
	}, nil
}
