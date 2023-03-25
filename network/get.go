package network

import (
	"github.com/austinkline/airdrop/db"
	"github.com/austinkline/airdrop/types"
)

const (
	queryGetNetwork     = `SELECT name, rpc_url FROM network WHERE name = ?;`
	queryGetAllNetworks = `SELECT name, rpc_url FROM network;`
)

// Get - retrieves a network from the database
// by name
func Get(name string) (n types.Network, err error) {
	c, err := db.GetConnection()
	if err != nil {
		return
	}

	defer c.Close()
	err = c.QueryRow(queryGetNetwork, name).Scan(&n.Name, &n.RpcURL)
	return
}

// GetAll - retrieves all networks from the database
func GetAll() (networks []types.Network, err error) {
	c, err := db.GetConnection()
	if err != nil {
		return
	}

	defer c.Close()
	rows, err := c.Query(queryGetAllNetworks)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var n types.Network
		err = rows.Scan(&n.Name, &n.RpcURL)
		if err != nil {
			return
		}

		networks = append(networks, n)
	}
	return
}
