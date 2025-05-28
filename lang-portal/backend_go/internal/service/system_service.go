package service

import (
	"database/sql"
	"errors"
	"time"

	"github.com/erans/lang-portal/internal/models"
)

// SystemService handles system-wide operations
type SystemService struct {
	db *sql.DB
}

// NewSystemService creates a new SystemService
func NewSystemService(db *sql.DB) *SystemService {
	return &SystemService{db: db}
}

// GetSystemStats retrieves system-wide statistics
func (s *SystemService) GetSystemStats() (*models.SystemStats, error) {
	var stats models.SystemStats

	// Get total words count
	err := s.db.QueryRow("SELECT COUNT(*) FROM words").Scan(&stats.TotalWords)
	if err != nil {
		return nil, err
	}

	// Get total groups count
	err = s.db.QueryRow("SELECT COUNT(*) FROM groups").Scan(&stats.TotalGroups)
	if err != nil {
		return nil, err
	}

	// Get total study sessions count
	err = s.db.QueryRow("SELECT COUNT(*) FROM study_sessions").Scan(&stats.TotalSessions)
	if err != nil {
		return nil, err
	}

	// Get average session score
	err = s.db.QueryRow(`
		SELECT COALESCE(AVG(score), 0) 
		FROM study_sessions 
		WHERE status = 'completed'
	`).Scan(&stats.AverageSessionScore)
	if err != nil {
		return nil, err
	}

	// Get total study time in minutes
	err = s.db.QueryRow(`
		SELECT COALESCE(SUM(CAST(
			(strftime('%s', end_time) - strftime('%s', start_time)) AS INTEGER
		) / 60), 0)
		FROM study_sessions 
		WHERE status = 'completed'
	`).Scan(&stats.TotalStudyTimeMinutes)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

// GetSystemHealth checks the system's health
func (s *SystemService) GetSystemHealth() (*models.SystemHealth, error) {
	health := &models.SystemHealth{
		Status:    "healthy",
		Timestamp: time.Now(),
	}

	// Check database connection
	err := s.db.Ping()
	if err != nil {
		health.Status = "unhealthy"
		health.Message = "Database connection error: " + err.Error()
		return health, nil
	}

	// Check if essential tables exist
	tables := []string{"words", "groups", "word_groups", "study_sessions", "study_activities"}
	for _, table := range tables {
		var count int
		err := s.db.QueryRow("SELECT COUNT(*) FROM " + table + " LIMIT 1").Scan(&count)
		if err != nil {
			health.Status = "unhealthy"
			health.Message = "Table check error for " + table + ": " + err.Error()
			return health, nil
		}
	}

	return health, nil
}

// BackupDatabase creates a backup of the database
func (s *SystemService) BackupDatabase(backupPath string) error {
	if backupPath == "" {
		return errors.New("backup path is required")
	}

	// For SQLite, we can use the backup API or simply copy the database file
	_, err := s.db.Exec("VACUUM INTO ?", backupPath)
	if err != nil {
		return err
	}

	return nil
}

// GetDatabaseSize returns the size of the database in bytes
func (s *SystemService) GetDatabaseSize() (int64, error) {
	var pageCount, pageSize int64

	err := s.db.QueryRow("PRAGMA page_count").Scan(&pageCount)
	if err != nil {
		return 0, err
	}

	err = s.db.QueryRow("PRAGMA page_size").Scan(&pageSize)
	if err != nil {
		return 0, err
	}

	return pageCount * pageSize, nil
}

// GetLastBackupInfo retrieves information about the last database backup
func (s *SystemService) GetLastBackupInfo() (*models.BackupInfo, error) {
	var info models.BackupInfo

	err := s.db.QueryRow(`
		SELECT backup_path, created_at, size_bytes
		FROM backup_history
		ORDER BY created_at DESC
		LIMIT 1
	`).Scan(&info.Path, &info.CreatedAt, &info.SizeBytes)

	if err == sql.ErrNoRows {
		return nil, errors.New("no backup history found")
	}
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// PruneOldData removes old data based on retention policy
func (s *SystemService) PruneOldData(retentionDays int) error {
	if retentionDays <= 0 {
		return errors.New("retention days must be positive")
	}

	// Begin transaction
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Delete old study sessions
	cutoffDate := time.Now().AddDate(0, 0, -retentionDays)
	_, err = tx.Exec(`
		DELETE FROM study_sessions
		WHERE end_time < ? AND status = 'completed'
	`, cutoffDate)
	if err != nil {
		return err
	}

	// Delete orphaned word review items
	_, err = tx.Exec(`
		DELETE FROM word_review_items
		WHERE session_id NOT IN (SELECT id FROM study_sessions)
	`)
	if err != nil {
		return err
	}

	// Commit transaction
	return tx.Commit()
}
