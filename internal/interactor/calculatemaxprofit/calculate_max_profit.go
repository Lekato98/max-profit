package calculatemaxprofit

import "context"

func (itr *Interactor) CalculateMaxProfit(ctx context.Context, profits []int64) (int64, error) {
	if len(profits) == 0 {
		return 0, nil
	}

	return profits[0], nil
}
