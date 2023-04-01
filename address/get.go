package address

import "github.com/austinkline/airdrop/db"

const (
	queryGetAccountByName = `
		SELECT addr, pk, name
		FROM address
		WHERE name = ?;`
	queryGetAccountByAddress = `
		SELECT addr, pk, name
		FROM address
		WHERE addr = ?;`
)

type Account struct {
	Address string `json:"address"`
	PK      string `json:"pk"`
	Name    string `json:"name"`
}

func GetAccountByName(name string) (account Account, err error) {
	// get a connection to the database
	db, err := db.GetConnection()
	if err != nil {
		return
	}

	// close the connection when we're done
	defer db.Close()

	// get the account
	err = db.QueryRow(queryGetAccountByName, name).Scan(&account.Address, &account.PK, &account.Name)
	return
}

func GetAccountByAddress(addr string) (account Account, err error) {
	// get a connection to the database
	db, err := db.GetConnection()
	if err != nil {
		return
	}

	// close the connection when we're done
	defer db.Close()

	// get the account
	err = db.QueryRow(queryGetAccountByAddress, addr).Scan(&account.Address, &account.PK, &account.Name)
	return
}
