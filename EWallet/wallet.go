package EWallet

type Wallet struct {
	id      int
	Address string `json:"address" binding:"required"`
	Amount  string `json:"amount" binding:"required"`
}
