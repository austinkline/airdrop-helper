package coinbase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendEth(t *testing.T) {
	destination := "0x10bF3f7Bc9Fba8354f46ec7a5588D87f1dE0896d"
	transaction, err := SendEth(destination, "0.1")
	assert.Nil(t, err)
	assert.Equal(t, destination, transaction.To.Address)
}
