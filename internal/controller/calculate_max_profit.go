package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CalculateMaxProfitHandler(c *gin.Context) {
	requestPayload := RequestPayloadDTO{}
	if err := c.ShouldBindBodyWithJSON(&requestPayload); err != nil {
		_ = c.Error(err)
	}

	if len(requestPayload) == 0 {
		_ = c.Error(fmt.Errorf("invalid paylad: empty array is not allowed"))
	}

	c.JSON(200, MaxProfitResponseDTO{
		Result: requestPayload[0],
	})
}
