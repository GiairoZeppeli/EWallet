package ewallet

type Transaction struct {
	Id     int     `json:"id"`
	From   string  `json:"from" binding:"required" db:"wallet_from"`
	To     string  `json:"to" binding:"required" db:"wallet_to"`
	Amount float32 `json:"amount" binding:"required" db:"amount"`
}

type WalletsTransactions struct {
	Id            int
	WalletId      int
	TransactionId int
}
