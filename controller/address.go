package controller

import (
	"challenge/model"
	"challenge/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDetails(c *gin.Context) {
	address := c.Param("address")

	fmt.Println(address)

	resp, err := http.Get("https://blockbook-bitcoin.tronwallet.me/api/v2/address/" + address)

	if err != nil {
		utils.Errors(http.StatusBadRequest, "Algo de errado não está certo.", err, c)
	}

	defer resp.Body.Close()

	var resposta model.BlockbookAddressDetails

	if err := json.NewDecoder(resp.Body).Decode(&resposta); err != nil {
		utils.Errors(http.StatusBadRequest, "Erro no decoder", err, c)
	}

	balance, err := strconv.Atoi(resposta.Balance)
	if err != nil {
		utils.Errors(http.StatusBadRequest, "Erro no parse - Balance", err, c)
	}

	unconfirmed, err := strconv.Atoi(resposta.UnconfirmedBalance)
	if err != nil {
		utils.Errors(http.StatusBadRequest, "Erro no parse - Unconfirmed", err, c)
	}

	confirmed := strconv.Itoa(balance - unconfirmed)

	c.JSON(http.StatusOK, &model.Wallet{
		Address: address,
		TotalTx: resposta.Txs,
		Balance: model.Balance{
			Confirmed:   confirmed,
			Unconfirmed: strconv.Itoa(unconfirmed),
		},
		Total: model.Total{
			Sent:    resposta.TotalSent,
			Receive: resposta.TotalReceived,
		},
	})
}
