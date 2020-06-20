package customer

import (
	"fmt"
	"github.com/pasuts/goonline/database"
)

func createTable() error {
	createTableStatement := `
	CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);
	`

	_, err := database.Conn().Exec(createTableStatement)
	if err != nil {
		return fmt.Errorf("cannot create database error : %w", err)
	}

	return nil
}