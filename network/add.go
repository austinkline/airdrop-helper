package network

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/austinkline/airdrop/db"
	"github.com/austinkline/airdrop/types"
)

const (
	queryAddNetwork = `
		INSERT INTO network
			(name, rpc_url)
		VALUES
			(?, ?)
		ON DUPLICATE KEY UPDATE
			rpc_url = VALUES(rpc_url);`
)

var ErrNetworkNotCreated = fmt.Errorf("network not created")

func Add(network types.Network) error {
	c, err := db.GetConnection()
	if err != nil {
		return err
	}

	defer c.Close()
	res, err := c.Exec(queryAddNetwork, network.Name, network.RpcURL)
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()
	if err != nil {
		log.WithError(err).Error("failed to get rows affected")
		return err
	}

	// the number of rows affected should be 1 since we added a row
	if ra < 1 {
		return ErrNetworkNotCreated
	}

	return nil
}
