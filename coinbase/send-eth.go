package coinbase

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	transactionsURL = "https://api.coinbase.com/v2/accounts/%s/transactions"
)

func SendEth(to, amount string) (transaction Transaction, err error) {
	payload := sendPayload{
		Type:     "send",
		To:       to,
		Amount:   amount,
		Currency: "ETH",
	}

	// serialize payload
	b, _ := json.Marshal(payload)
	reader := bytes.NewReader(b)

	url := fmt.Sprintf(transactionsURL, ethAccountID)
	req, _ := http.NewRequest("POST", url, reader)
	req.Header.Set("Content-Type", "application/json")

	// sign request
	err = signRequest(req)
	if err != nil {
		return
	}

	// send request
	c := http.Client{}
	response, err := c.Do(req)
	if err != nil {
		return
	}

	var cbResponse Response[Transaction]
	responseBytes, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(responseBytes, &cbResponse)
	if err != nil {
		log.Warn(string(responseBytes))
		log.WithError(err).Error("Failed to unmarshal response")
		return
	}

	transaction = cbResponse.Data
	return
}

type sendPayload struct {
	Type        string  `json:"type"` // must be "send"
	To          string  `json:"to"`
	Amount      string  `json:"amount"`
	Currency    string  `json:"currency"`
	Description *string `json:"description,omitempty"`
	Idem        *string `json:"idem,omitempty"`
	// For select currencies, destination_tag or memo indicates the beneficiary or destination of a
	// payment for select currencies. Example:
	// { "type" : "send", "to": "address", "destination_tag" : "memo", "amount": "", "currency": "" }
	DestinationTag *string `json:"destination_tag,omitempty"`
}

type Transaction struct {
	Id           string    `json:"id"`
	Type         string    `json:"type"`
	Status       string    `json:"status"`
	Amount       Amount    `json:"amount"`
	NativeAmount Amount    `json:"native_amount"`
	Description  *string   `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Resource     string    `json:"resource"`
	ResourcePath string    `json:"resource_path"`
	Network      Network   `json:"network"`
	To           Receiver  `json:"to"`
	Details      Details   `json:"details"`
}

type Amount struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Network struct {
	Status string `json:"status"`
	Hash   string `json:"hash"`
	Name   string `json:"name"`
}

type Receiver struct {
	Resource string `json:"resource"`
	Address  string `json:"address"`
}

type Details struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}
