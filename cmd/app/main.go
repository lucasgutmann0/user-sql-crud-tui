/*
Copyright © 2024 Lucas Gutmann email@example.com
*/
package main

import (
	"os"

	"github.com/lucasgutmann0/user-sql-crud-tui/internal/cmd"
	"github.com/lucasgutmann0/user-sql-crud-tui/internal/database"
)

func main() {
	// Create database
	dbName := "user-database.sql"
	db, err := database.Setup(&dbName)
	if err != nil {
		os.Exit(1)
	}
	defer db.Close()

	// Make migrations to the DB
	database.MakeMigrations(db, database.DatabaseSchema)

	cmd.Execute()
}
