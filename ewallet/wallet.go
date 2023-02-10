package ewallet

type Wallet struct {
	Id      int     `json:"id"`
	Address string  `json:"address" binding:"required"  db:"address"`
	Amount  float32 `json:"amount" binding:"required" db:"amount"`
}
