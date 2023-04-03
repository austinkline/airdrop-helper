package coinbase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	listTransactionsURL = "https://api.coinbase.com/v2/accounts/%s/transactions"
)

func GetTransactions() (transactions []Transaction, err error) {
	url := fmt.Sprintf(listTransactionsURL, ethAccountID)
	req, _ := http.NewRequest("GET", url, nil)
	err = signRequest(req)
	if err != nil {
		return
	}

	c := http.Client{}
	response, err := c.Do(req)
	if err != nil {
		return
	}

	// read body from response into a bytes array
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var cbResponse Response[[]Transaction]
	err = json.Unmarshal(bytes, &cbResponse)
	if err != nil {
		return
	}

	transactions = cbResponse.Data
	return
}
