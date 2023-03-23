package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	host, user, password, db string

	envHost     = "MYSQL_HOST"
	envUser     = "MYSQL_USER"
	envPassword = "MYSQL_PASSWORD"
	envDB       = "MYSQL_DB"
)

// GetConnection returns a new mysql connection
// to our database, reading the environment variables
// MYSQL_HOST, MYSQL_USER, MYSQL_PASSWORD, and MYSQL_DB
// to determine the connection parameters.
func GetConnection() (*sql.DB, error) {
	// username:password@tcp(127.0.0.1:3306)/dbname
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=true", user, password, host, "3306", db)
	db, err := sql.Open("mysql", connectionString)
	return db, err
}

func init() {
	host = os.Getenv(envHost)
	user = os.Getenv(envUser)
	password = os.Getenv(envPassword)
	db = os.Getenv(envDB)
}
