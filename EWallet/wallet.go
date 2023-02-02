package EWallet

type Wallet struct {
	Address string  `json:"address" binding:"required"`
	Amount  float32 `json:"amount" binding:"required"`
}
