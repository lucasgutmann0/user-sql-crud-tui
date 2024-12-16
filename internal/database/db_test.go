package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Cases
// 1. InMemoryDatabaseNilPath: Database created in memory (nil param)
// 2. InMemoryDatabaseWithPath: created in memory (":memory:")
// 3. FileDatabaseWithPath: Database created in file ("test.db.sql")
// 4. DatabaseWithBadPath: Function receives bad path
func TestDatabaseSetup(t *testing.T) {
	// 1. InMemoryDatabaseNilPath: Database created in memory (nil param)
	t.Run("InMemoryDatabaseNilPath", func(t *testing.T) {
		db, err := Setup(nil)
		require.NoError(t, err, "Should create in-memory database without "+
			"error and nil dbPath param")
		require.NotNil(t, db, "Database object should not be nil")
		defer db.Close()
	})

	// 2. InMemoryDatabaseWithPath: created in memory (":memory:")
	t.Run("InMemoryDatabaseWithPath", func(t *testing.T) {
		// AAA: Pattern - Arrange, Act, Assert
		// Arrange
		dbPath := ":memory:"

		// Act
		db, err := Setup(&dbPath)

		// Assert
		require.NoError(
			t, err, "Should create in-memory database without error")
		require.NotNil(t, db, "Database object should not be nil")

		// Cleanup
		defer db.Close()
	})

	// 3. FileDatabaseWithPath: Database created in file ("test.db.sql")
	t.Run("FileDatabaseWithPath", func(t *testing.T) {
		dbPath := "test.db.sql"
		db, err := Setup(&dbPath)
		require.NoErrorf(t, err, "Should create in file database without "+
			" error in path: %s\n", dbPath)
		require.NotNil(t, db, "Database object should not be nil")

		// Clean up
		defer db.Close()
		defer cleanupCreatedFile(t, dbPath)
	})

	// 4. DatabaseWithBadPath: Function receives bad path
	t.Run("DatabaseWithBadPath", func(t *testing.T) {
		dbPath := "C:/failing-purpousely/?.sql" // Bad path intentionally
		db, err := Setup(&dbPath)
		require.Error(t, err, "Should have an error because of bad path")
		require.Nil(t, db, "Database object SHOULD be nil")
	})
}

func cleanupCreatedFile(t *testing.T, path string) {
	// Check if file exists
	_, err := os.Stat(path)
	require.NoError(t, err, "File should exist in file system")

	// Delete file if exists
	if err := os.Remove(path); err != nil {
		require.NoError(t, err, "File should have been removed successfully")
	}
}
