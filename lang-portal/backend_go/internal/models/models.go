package models

import (
	"time"
)

// Word represents a vocabulary word in the system
type Word struct {
	ID       int64          `json:"id"`
	Japanese string         `json:"japanese"`
	Romaji   string         `json:"romaji"`
	English  string         `json:"english"`
	Parts    map[string]any `json:"parts"`
}

// Group represents a collection of words
type Group struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// WordGroup represents the many-to-many relationship between words and groups
type WordGroup struct {
	ID      int64 `json:"id"`
	WordID  int64 `json:"word_id"`
	GroupID int64 `json:"group_id"`
}

// StudySession represents a learning session
type StudySession struct {
	ID              int64      `json:"id"`
	StartTime       time.Time  `json:"start_time"`
	EndTime         *time.Time `json:"end_time,omitempty"`
	Score           *float64   `json:"score,omitempty"`
	Status          string     `json:"status"` // active, completed, abandoned
	StudyActivityID int64      `json:"study_activity_id"`
}

// StudyActivity represents a study activity type
type StudyActivity struct {
	ID           int64     `json:"id"`
	GroupID      int64     `json:"group_id"`
	ActivityType string    `json:"activity_type"`
	CreatedAt    time.Time `json:"created_at"`
}

// WordReviewItem represents a single word review instance
type WordReviewItem struct {
	ID         int64     `json:"id"`
	SessionID  int64     `json:"session_id"`
	WordID     int64     `json:"word_id"`
	IsCorrect  bool      `json:"is_correct"`
	ReviewedAt time.Time `json:"reviewed_at"`
}
