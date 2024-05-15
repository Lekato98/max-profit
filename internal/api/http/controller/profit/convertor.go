package profit

func toMaxProfitResponseDTO(maxProfit int64) MaxProfitResponseDTO {
	return MaxProfitResponseDTO{
		Result: maxProfit,
	}
}
