package coinbase

import "testing"

func TestGetUser(t *testing.T) {
	_, err := GetUser()
	if err != nil {
		t.Error(err)
	}
}
