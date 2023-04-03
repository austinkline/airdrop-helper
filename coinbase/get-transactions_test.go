package coinbase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTransactions(t *testing.T) {
	_, err := GetTransactions()
	assert.Nil(t, err)
}
