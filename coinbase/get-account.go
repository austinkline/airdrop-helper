package coinbase

import (
	"encoding/json"
	"github.com/austinkline/airdrop/db"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	listAccountsURL = "https://api.coinbase.com/v2/accounts"

	queryAccountBySymbol = "SELECT account_id FROM cb_account WHERE symbol = ? LIMIT 1"
)

var (
	ethAccountID = os.Getenv("CB_ETH_ACCOUNT_ID")
)

func GetIdForSymbol(symbol string) (id string, err error) {
	// query our mysql database for the id
	// of the account with the given symbol

	// if there are multiple accounts with the same symbol,
	// return the first one

	// get a connection to our database
	connection, err := db.GetConnection()
	if err != nil {
		return
	}

	// query the database
	rows, err := connection.Query(queryAccountBySymbol, symbol)
	if err != nil {
		return
	}

	// iterate over the rows
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return
		}
	}

	return
}

func GetAccount(id string) (account Account, err error) {
	url := listAccountsURL + "/" + id
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

func GetEthAccount() (account Account, err error) {
	return GetAccount(ethAccountID)
}

func GetAccountForSymbol(symbol string) (account Account, err error) {
	id, err := GetIdForSymbol(symbol)
	if err != nil {
		return
	}

	account, err = GetAccount(id)
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
