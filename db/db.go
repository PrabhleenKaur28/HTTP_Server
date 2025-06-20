package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"//(blank import) helper function to run services of database/sql. This repo is not used directly hence is imported with a _ else it will cause an error for imported but not used
)

var DB *sql.DB //creates a global variable named DB that holds connection to PostgreSQL database 
// sql.DB is a type from Go's database/sql package that represents a database connection pool
//* means it's a pointer so DB points to the actual connection pool in memory
//*sql.DB is thread-safe as it manages connections automatically

func Init() error {//return type is error because connecting to database is an operation that can fail
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",//no encryption is needed so sslmode is disabled because we are working on localhost for now
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = sql.Open("postgres", connStr)//sql.Open() prepares the connection
	if err != nil {
		return err
	}

	return DB.Ping()//Ping() actually verifies connection works
}
