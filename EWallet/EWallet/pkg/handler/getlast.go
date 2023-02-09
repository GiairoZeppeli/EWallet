package handler

import (
	"github.com/GiairoZeppeli/EWallet"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
)

func (h *Handler) GetLast(c *gin.Context) {
	u, err := url.Parse(c.FullPath())
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	q := u.Query()
	countString := q["count"]
	number, err := strconv.ParseUint(countString[0], 10, 32)
	count := int(number)

	transactions, err := h.services.GetLast.GetLast(count)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getTransactionsResponce{
		Data: transactions,
	})
}

type getTransactionsResponce struct {
	Data []EWallet.Transaction `json:"data"`
}
