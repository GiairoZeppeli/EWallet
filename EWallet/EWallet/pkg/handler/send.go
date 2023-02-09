package handler

import (
	"github.com/GiairoZeppeli/EWallet"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Send(c *gin.Context) {
	var input EWallet.Transaction

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.Send.Send(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
