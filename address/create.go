package address

import "os"

const (
	envMnemonic = "MNEMONIC"

	queryInsertAddress = `
		INSERT INTO address (address, name, pk)
		VALUES ($1, $2, $3);`
)

var (
	mnemonic string
)

// Create creates a new ethereum address, using the mneumonic variable as the seed
// for the key generation
// Then inserts the address into our database with the given name,
// so it can be referenced later.
func Create(name string) (string, error) {
	return "", nil
}

func init() {
	mnemonic = os.Getenv(envMnemonic)
}
