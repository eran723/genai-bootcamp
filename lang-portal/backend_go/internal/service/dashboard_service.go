package service

import (
	"database/sql"
	"time"
)

// DashboardService handles dashboard-related business logic
type DashboardService struct {
	db *sql.DB
}

// NewDashboardService creates a new DashboardService
func NewDashboardService(db *sql.DB) *DashboardService {
	return &DashboardService{db: db}
}

// LastSessionResponse represents the last study session details
type LastSessionResponse struct {
	SessionID      int64      `json:"session_id"`
	StartTime      time.Time  `json:"start_time"`
	EndTime        *time.Time `json:"end_time,omitempty"`
	Score          float64    `json:"score"`
	Status         string     `json:"status"`
	ActivityType   string     `json:"activity_type"`
	GroupID        int64      `json:"group_id"`
	GroupName      string     `json:"group_name"`
	WordsReviewed  int        `json:"words_reviewed"`
	CorrectAnswers int        `json:"correct_answers"`
}

// StatsResponse represents study statistics
type StatsResponse struct {
	TotalStudyTime     int     `json:"total_study_time"` // in seconds
	SessionsCompleted  int     `json:"sessions_completed"`
	TotalWordsReviewed int     `json:"total_words_reviewed"`
	SuccessRate        float64 `json:"success_rate"`
	StudyStreakDays    int     `json:"study_streak_days"`
}

// ProgressResponse represents learning progress
type ProgressResponse struct {
	OverallCompletion   float64 `json:"overall_completion"`
	TotalWordsStudied   int     `json:"total_words_studied"`
	TotalAvailableWords int     `json:"total_available_words"`
}

// GetLastSession retrieves details about the most recent study session
func (s *DashboardService) GetLastSession() (*LastSessionResponse, error) {
	query := `
		SELECT 
			ss.id,
			ss.start_time,
			ss.end_time,
			ss.score,
			ss.status,
			sa.activity_type,
			g.id,
			g.name,
			(SELECT COUNT(*) FROM word_review_items wri WHERE wri.activity_id = sa.id) as words_reviewed,
			(SELECT COUNT(*) FROM word_review_items wri WHERE wri.activity_id = sa.id AND wri.is_correct = true) as correct_answers
		FROM study_sessions ss
		JOIN study_activities sa ON sa.session_id = ss.id
		JOIN groups g ON g.id = sa.group_id
		ORDER BY ss.start_time DESC
		LIMIT 1
	`

	var resp LastSessionResponse
	err := s.db.QueryRow(query).Scan(
		&resp.SessionID,
		&resp.StartTime,
		&resp.EndTime,
		&resp.Score,
		&resp.Status,
		&resp.ActivityType,
		&resp.GroupID,
		&resp.GroupName,
		&resp.WordsReviewed,
		&resp.CorrectAnswers,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No sessions yet
		}
		return nil, err
	}

	return &resp, nil
}

// GetStats retrieves study statistics
func (s *DashboardService) GetStats() (*StatsResponse, error) {
	query := `
		WITH daily_sessions AS (
			SELECT 
				DATE(start_time) as study_date,
				COUNT(*) as sessions_per_day
			FROM study_sessions
			GROUP BY DATE(start_time)
		),
		streak AS (
			SELECT COUNT(*) as streak_days
			FROM (
				SELECT study_date,
				       ROW_NUMBER() OVER (ORDER BY study_date DESC) as rn,
				       DATE(study_date, '-' || ROW_NUMBER() OVER (ORDER BY study_date DESC) || ' days') as grp
				FROM daily_sessions
			)
			WHERE grp = DATE('now', '-' || rn || ' days')
		)
		SELECT
			COALESCE(SUM(CASE 
				WHEN end_time IS NOT NULL THEN 
					ROUND((JULIANDAY(end_time) - JULIANDAY(start_time)) * 86400)
				ELSE 0 
			END), 0) as total_study_time,
			COUNT(*) FILTER (WHERE status = 'completed') as completed_sessions,
			(SELECT COUNT(*) FROM word_review_items) as total_reviews,
			COALESCE(AVG(CASE WHEN is_correct THEN 100.0 ELSE 0.0 END), 0) as success_rate,
			COALESCE((SELECT streak_days FROM streak), 0) as streak_days
		FROM study_sessions ss
		LEFT JOIN study_activities sa ON sa.session_id = ss.id
		LEFT JOIN word_review_items wri ON wri.activity_id = sa.id
	`

	var stats StatsResponse
	err := s.db.QueryRow(query).Scan(
		&stats.TotalStudyTime,
		&stats.SessionsCompleted,
		&stats.TotalWordsReviewed,
		&stats.SuccessRate,
		&stats.StudyStreakDays,
	)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

// GetProgress retrieves learning progress information
func (s *DashboardService) GetProgress() (*ProgressResponse, error) {
	query := `
		WITH studied_words AS (
			SELECT DISTINCT word_id
			FROM word_review_items
		)
		SELECT
			CASE 
				WHEN COUNT(w.id) > 0 THEN 
					ROUND(COUNT(DISTINCT sw.word_id) * 100.0 / COUNT(w.id), 1)
				ELSE 0.0
			END as completion_rate,
			COUNT(DISTINCT sw.word_id) as words_studied,
			COUNT(w.id) as total_words
		FROM words w
		LEFT JOIN studied_words sw ON sw.word_id = w.id
	`

	var progress ProgressResponse
	err := s.db.QueryRow(query).Scan(
		&progress.OverallCompletion,
		&progress.TotalWordsStudied,
		&progress.TotalAvailableWords,
	)
	if err != nil {
		return nil, err
	}

	return &progress, nil
}
