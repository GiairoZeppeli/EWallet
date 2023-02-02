package EWallet

type Transaction struct {
	From   string  `json:"from" binding:"required"`
	To     string  `json:"to" binding:"required"`
	Amount float32 `json:"amount" binding:"required"`
}
