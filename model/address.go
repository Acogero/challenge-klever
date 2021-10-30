package model

// func Address() {
// address,
// balance,
// totalTx, - txs
// balance, - { confirmed (balance - unconfirmed), unconfirmed, }
// "total": {
// "sent": "1189163719343",
// "received": "1189307731097"
// }
// }

type Wallet struct {
	Address string  `json:"address"`
	TotalTx int64   `json:"totalTx"`
	Balance Balance `json:"balance"`
	Total   Total   `json:"total"`
}

type Total struct {
	Sent    string `json:"sent"`
	Receive string `json:"receive"`
}

type Balance struct {
	Confirmed   string `json:"confirmed"`
	Unconfirmed string `json:"unconfirmed"`
}
