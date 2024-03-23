package test

import (
	"testing"

	"x-clone.com/backend/src/database"
)

func TestConnectDB(t *testing.T) {
	// Provide test database credentials
	username := "postgres"
	password := "saurav"
	dbName := "x-clone"

	// Attempt to connect to the test database
	db, err := database.ConnectDB(username, password, dbName)
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}

	// Close the database connection
	defer db.Close()

	// Verify that the database connection is still alive
	if err := db.DB.Ping(); err != nil {
		t.Errorf("Error pinging database: %v", err)
	}
}
