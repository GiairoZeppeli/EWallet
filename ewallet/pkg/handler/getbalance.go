package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetBalance(c *gin.Context) {
	address := c.Param("address")
	var balance float32
	balance, err := h.services.GetBalance.GetBalance(address)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getBalanceResponse{
		Balance: balance})
}

type getBalanceResponse struct {
	Balance float32 `json:"balance"`
}
