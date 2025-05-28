//go:build mage
// +build mage

package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	_ "github.com/mattn/go-sqlite3"
)

// Default target to run when none is specified
var Default = Build

// Build builds the application
func Build() error {
	fmt.Println("Building...")
	return sh.Run("go", "build", "-o", "bin/server", "./cmd/server")
}

// InitDB initializes the SQLite database
func InitDB() error {
	fmt.Println("Initializing database...")

	dbPath := "words.db"
	if _, err := os.Stat(dbPath); err == nil {
		fmt.Println("Database already exists")
		return nil
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	return nil
}

// Migrate runs database migrations
func Migrate() error {
	mg.Deps(InitDB)

	fmt.Println("Running migrations...")

	db, err := sql.Open("sqlite3", "words.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	migrations, err := filepath.Glob("db/migrations/*.sql")
	if err != nil {
		return fmt.Errorf("failed to list migrations: %v", err)
	}

	for _, migration := range migrations {
		fmt.Printf("Applying migration: %s\n", migration)

		content, err := os.ReadFile(migration)
		if err != nil {
			return fmt.Errorf("failed to read migration %s: %v", migration, err)
		}

		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("failed to apply migration %s: %v", migration, err)
		}
	}

	return nil
}

// Seed adds sample data to the database
func Seed() error {
	mg.Deps(Migrate)

	fmt.Println("Seeding database...")

	db, err := sql.Open("sqlite3", "words.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	seeds, err := filepath.Glob("db/seeds/*.sql")
	if err != nil {
		return fmt.Errorf("failed to list seeds: %v", err)
	}

	for _, seed := range seeds {
		fmt.Printf("Applying seed: %s\n", seed)

		content, err := os.ReadFile(seed)
		if err != nil {
			return fmt.Errorf("failed to read seed %s: %v", seed, err)
		}

		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("failed to apply seed %s: %v", seed, err)
		}
	}

	return nil
}

// Clean removes generated files
func Clean() error {
	fmt.Println("Cleaning...")
	os.Remove("words.db")
	return nil
}

// Dev runs the application in development mode
func Dev() error {
	mg.Deps(Migrate)
	fmt.Println("Starting server in development mode...")
	return sh.Run("go", "run", "./cmd/server")
}
