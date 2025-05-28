package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Initialize initializes the database connection
func Initialize() error {
	// Ensure the database directory exists
	dbPath := filepath.Join(".", "words.db")

	// For development, remove the database if it exists
	if err := os.Remove(dbPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove existing database: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return err
	}

	// Open the database
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return err
	}

	// Enable foreign keys
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	log.Println("Database connection established")
	return nil
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return db
}

// Close closes the database connection
func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// RunMigrations runs all migration files in order
func RunMigrations() error {
	files, err := filepath.Glob("db/migrations/*.sql")
	if err != nil {
		return fmt.Errorf("failed to list migration files: %w", err)
	}

	log.Println("Running migrations...")
	for _, file := range files {
		log.Printf("Executing migration: %s", filepath.Base(file))
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", file, err)
		}

		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", file, err)
		}
	}

	log.Println("Migrations completed successfully")
	return nil
}

// RunSeeds runs all seed files
func RunSeeds() error {
	// TODO: Implement seeding from JSON files
	return nil
}
