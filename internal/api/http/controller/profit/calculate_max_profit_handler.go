package profit

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctr *Controller) MaxProfitHandler(c *gin.Context) {
	requestPayload := RequestPayloadDTO{}
	if err := c.ShouldBindBodyWithJSON(&requestPayload); err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "invalid payload")
		return
	}

	if len(requestPayload) == 0 {
		c.String(http.StatusBadRequest, "empty array is not allowed")
		return
	}

	result, err := ctr.calculateMaxProfitInteractor.CalculateMaxProfit(c.Request.Context(), requestPayload)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, toMaxProfitResponseDTO(result))
}
