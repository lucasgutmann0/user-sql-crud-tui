package database

import (
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/lucasgutmann0/user-sql-crud-tui/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

// Setup initializes the database with the given path (dbPath) name.
// Use ":memory:" or nil for an in-memory database during testing.
func Setup(dbPath *string) (*sqlx.DB, error) {
	// Use defaults
	if dbPath == nil {
		slog.Warn("Database path wasn't defined. Using default in memory db")
		defaultPath := ":memory:"
		dbPath = &defaultPath
	}

	// Create connection to db
	db, err := sqlx.Connect("sqlite3", *dbPath)
	if err != nil {
		slog.Error("Failed to start connection or open the DB", "error", err,
			"method", "SetUpDatabase")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		slog.Error("Failed at pinging the created DB", "error", err,
			"method", "SetUpDatabase")
		return nil, err
	}

	return db, nil
}

// Migrate the schema to sqlite db
func MakeMigrations(db *sqlx.DB, migrationSchema string) error {
	slog.Info("Starting migration to apply predefined schemas.")

	// Create database tables and check for errors
	if _, err := db.Exec(migrationSchema); err != nil {
		slog.Error("Schema migration failed", "error", err,
			"method", "MakeMigrations")
		return err
	}

	slog.Info("Schema migration was successful")
	return nil
}

func SeedDatase(db *sqlx.DB, seedData []models.User) error {
	result, err := db.Exec(`
		INSERT INTO users (id, name, role, email, age, password)
		VALUES (:id, :name, :role, :email, :age, :password)`, seedData)
	if err != nil {
		slog.Error("Database seeding query failed",
			"error", err, "method", "SeedDatabase")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("Database seeding failed", "error", err,
			"method", "SeedDatabase")
		return err
	}

	msg := fmt.Sprintf("Database Seeding was successful. "+
		"Total rows inserted: %d.", rowsAffected)
	slog.Info(msg)

	return nil
}
