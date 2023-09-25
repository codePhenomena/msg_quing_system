package models

import (
    "testing"
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "os"
)

const (
    DB_USER = "root"
    DB_PASSWORD = "Root@123"
    DB_NAME = "product_details" 
)

var testDB *sqlx.DB // Use a DB connection variable for tests

func TestMain(m *testing.M) {
    db, err := sqlx.Open("mysql", DB_USER+":"+DB_PASSWORD+"@/"+DB_NAME)
    if err != nil {
        log.Fatal("cannot connect to db", err)
    }
    
    testDB = db // Initialize testDB with the database connection
    
    // Run the tests
    exitCode := m.Run()

    // Close the testDB connection
    if err := testDB.Close(); err != nil {
        log.Fatal("error closing db connection", err)
    }

    // Exit with the appropriate exit code
    os.Exit(exitCode)
}


