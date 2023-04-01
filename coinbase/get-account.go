package coinbase

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	listAccountsURL = "https://api.coinbase.com/v2/accounts"
)

var (
	ethAccountID = os.Getenv("CB_ETH_ACCOUNT_ID")
)

func GetEthAccount() (account Account, err error) {
	url := listAccountsURL + "/" + ethAccountID
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

	var cbResponse Response[Account]

	// unmarshal the bytes array into a map
	err = json.Unmarshal(bytes, &cbResponse)
	if err != nil {
		return
	}

	account = cbResponse.Data
	return
}

func GetAccountForSymbol(symbol string) (account Account, err error) {
	accounts, err := GetAccounts()
	if err != nil {
		return
	}

	for _, a := range accounts {
		if a.Currency.Code == symbol {
			account = a
			return
		}
	}

	return
}

func GetAccounts() (accounts []Account, err error) {
	req, _ := http.NewRequest("GET", listAccountsURL, nil)
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

	var cbResponse Response[[]Account]
	err = json.Unmarshal(bytes, &cbResponse)
	if err != nil {
		return
	}

	accounts = cbResponse.Data
	return
}

type Account struct {
	Id               string    `json:"id"`
	Name             string    `json:"name"`
	Primary          bool      `json:"primary"`
	Type             string    `json:"type"`
	Currency         Currency  `json:"currency"`
	Balance          Balance   `json:"balance"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Resource         string    `json:"resource"`
	ResourcePath     string    `json:"resource_path"`
	AllowDeposits    bool      `json:"allow_deposits"`
	AllowWithdrawals bool      `json:"allow_withdrawals"`
}

type Currency struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	SortIndex    int    `json:"sort_index"`
	Exponent     int    `json:"exponent"`
	Type         string `json:"type"`
	AddressRegex string `json:"address_regex"`
	AssetId      string `json:"asset_id"`
	Slug         string `json:"slug"`
}

type Balance struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}
