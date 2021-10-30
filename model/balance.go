package model

type Wallet struct {
	Address string  `json:"address"`
	TotalTx int64   `json:"totalTx"`
	Balance Balance `json:"balance"`
	Total   Total   `json:"total"`
}
