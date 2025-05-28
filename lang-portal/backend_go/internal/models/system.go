package models

import "time"

// SystemStats represents system-wide statistics
type SystemStats struct {
	TotalWords            int64   `json:"total_words"`
	TotalGroups           int64   `json:"total_groups"`
	TotalSessions         int64   `json:"total_sessions"`
	AverageSessionScore   float64 `json:"average_session_score"`
	TotalStudyTimeMinutes int64   `json:"total_study_time_minutes"`
}

// SystemHealth represents the system's health status
type SystemHealth struct {
	Status    string    `json:"status"`
	Message   string    `json:"message,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

// BackupInfo represents information about a database backup
type BackupInfo struct {
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
	SizeBytes int64     `json:"size_bytes"`
}
