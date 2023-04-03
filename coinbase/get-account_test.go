package coinbase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccounts(t *testing.T) {
	_, err := GetAccounts()
	assert.Nil(t, err)
}

func TestGetEthereumAccount(t *testing.T) {
	_, err := GetEthAccount()
	assert.Nil(t, err)
}

func TestGetAccountForSymbol(t *testing.T) {
	_, err := GetAccountForSymbol("eth")
	assert.Nil(t, err)
}
