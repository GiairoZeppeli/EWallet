package handler

import (
	"github.com/GiairoZeppeli/EWallet"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) Send(c *gin.Context) {
	var input EWallet.Transaction

	if err := c.BindJSON($input); err != nil{
		logrus.Errorf(),
	}
}
