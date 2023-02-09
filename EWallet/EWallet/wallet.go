package EWallet

type Wallet struct {
	Address string  `json:"address" binding:"required"  db:"address"`
	Amount  float32 `json:"amount" binding:"required" db:"amount"`
}
