package model

type BlockbookAddressDetails struct {
	Address            string `json:"address"`
	Balance            string `json:"balance"`
	TotalReceived      string `json:"totalReceived"`
	TotalSent          string `json:"totalSent"`
	UnconfirmedBalance string `json:"unconfirmedBalance"`
	UnconfirmedTxs     int32  `json:"unconfirmedTxs"`
	Txs                int64  `json:"txs"`
}
