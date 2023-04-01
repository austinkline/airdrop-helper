package address

import (
	"crypto/ecdsa"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"

	"github.com/austinkline/airdrop/db"
)

const (
	envMnemonic = "MNEMONIC"

	queryGetMaxAddressID = "SELECT MAX(id) FROM address"

	queryInsertAddress = `
		INSERT INTO address (addr, name, pk)
		VALUES (?, ?, ?);`
)

var (
	mnemonic string
)

func getLastID() (id int, err error) {
	// get sql connection
	db, err := db.GetConnection()
	if err != nil {
		return
	}

	// close connection when we're done
	defer db.Close()

	// get the max id
	err = db.QueryRow(queryGetMaxAddressID).Scan(&id)
	return
}

// Create a new ethereum address and store it in our database,
// so it can be referenced later. Returns the address of
// the private key and any errors.
func Create(name string) (address string, err error) {
	// get the last id
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	pk := hexutil.Encode(privateKeyBytes)[2:]
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = fmt.Errorf("error casting public key to ECDSA")
		return
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	address = hexutil.Encode(hash.Sum(nil)[12:])

	// get sql connection
	db, err := db.GetConnection()
	if err != nil {
		return
	}

	// close connection when we're done
	defer db.Close()

	// insert the address
	_, err = db.Exec(queryInsertAddress, address, name, pk)
	return
}

func init() {
	mnemonic = os.Getenv(envMnemonic)
}
