package service

import (
	"database/sql"
	"errors"
	"time"

	"github.com/erans/lang-portal/internal/models"
)

// StudySessionService handles business logic for study sessions
type StudySessionService struct {
	db *sql.DB
}

// NewStudySessionService creates a new StudySessionService
func NewStudySessionService(db *sql.DB) *StudySessionService {
	return &StudySessionService{db: db}
}

// GetSession retrieves a study session by ID
func (s *StudySessionService) GetSession(id int64) (*models.StudySession, error) {
	var session models.StudySession
	err := s.db.QueryRow(`
		SELECT id, start_time, end_time, score, status, study_activity_id
		FROM study_sessions 
		WHERE id = ?`,
		id,
	).Scan(
		&session.ID,
		&session.StartTime,
		&session.EndTime,
		&session.Score,
		&session.Status,
		&session.StudyActivityID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("study session not found")
		}
		return nil, err
	}

	return &session, nil
}

// ListSessions retrieves a paginated list of study sessions
func (s *StudySessionService) ListSessions(offset, limit int) ([]models.StudySession, error) {
	rows, err := s.db.Query(`
		SELECT id, start_time, end_time, score, status, study_activity_id
		FROM study_sessions 
		ORDER BY start_time DESC
		LIMIT ? OFFSET ?`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []models.StudySession
	for rows.Next() {
		var session models.StudySession
		if err := rows.Scan(
			&session.ID,
			&session.StartTime,
			&session.EndTime,
			&session.Score,
			&session.Status,
			&session.StudyActivityID,
		); err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}

	return sessions, nil
}

// CreateSession creates a new study session
func (s *StudySessionService) CreateSession(session *models.StudySession) error {
	session.StartTime = time.Now()
	session.Status = "in_progress"

	result, err := s.db.Exec(`
		INSERT INTO study_sessions (start_time, status, study_activity_id)
		VALUES (?, ?, ?)`,
		session.StartTime,
		session.Status,
		session.StudyActivityID,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	session.ID = id
	return nil
}

// UpdateSession updates an existing study session
func (s *StudySessionService) UpdateSession(session *models.StudySession) error {
	result, err := s.db.Exec(`
		UPDATE study_sessions 
		SET end_time = ?, score = ?, status = ?
		WHERE id = ?`,
		session.EndTime,
		session.Score,
		session.Status,
		session.ID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("study session not found")
	}

	return nil
}

// EndSession ends a study session and calculates the final score
func (s *StudySessionService) EndSession(id int64, score float64) error {
	result, err := s.db.Exec(`
		UPDATE study_sessions 
		SET end_time = ?, score = ?, status = 'completed'
		WHERE id = ?`,
		time.Now(),
		score,
		id,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("study session not found")
	}

	return nil
}

// GetSessionReviewItems retrieves all word review items for a session
func (s *StudySessionService) GetSessionReviewItems(sessionID int64) ([]models.WordReviewItem, error) {
	rows, err := s.db.Query(`
		SELECT id, word_id, session_id, is_correct, reviewed_at
		FROM word_review_items
		WHERE session_id = ?
		ORDER BY reviewed_at ASC`,
		sessionID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.WordReviewItem
	for rows.Next() {
		var item models.WordReviewItem
		if err := rows.Scan(
			&item.ID,
			&item.WordID,
			&item.SessionID,
			&item.IsCorrect,
			&item.ReviewedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
