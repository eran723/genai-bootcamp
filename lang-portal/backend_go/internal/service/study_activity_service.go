package service

import (
	"database/sql"
	"errors"
	"time"

	"github.com/erans/lang-portal/internal/models"
)

// StudyActivityService handles business logic for study activities
type StudyActivityService struct {
	db *sql.DB
}

// NewStudyActivityService creates a new StudyActivityService
func NewStudyActivityService(db *sql.DB) *StudyActivityService {
	return &StudyActivityService{db: db}
}

// GetActivity retrieves a study activity by ID
func (s *StudyActivityService) GetActivity(id int64) (*models.StudyActivity, error) {
	var activity models.StudyActivity
	err := s.db.QueryRow(`
		SELECT id, group_id, activity_type, created_at
		FROM study_activities 
		WHERE id = ?`,
		id,
	).Scan(
		&activity.ID,
		&activity.GroupID,
		&activity.ActivityType,
		&activity.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("study activity not found")
		}
		return nil, err
	}

	return &activity, nil
}

// ListActivities retrieves a paginated list of study activities
func (s *StudyActivityService) ListActivities(offset, limit int) ([]models.StudyActivity, error) {
	rows, err := s.db.Query(`
		SELECT id, group_id, activity_type, created_at
		FROM study_activities 
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []models.StudyActivity
	for rows.Next() {
		var activity models.StudyActivity
		if err := rows.Scan(
			&activity.ID,
			&activity.GroupID,
			&activity.ActivityType,
			&activity.CreatedAt,
		); err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}

	return activities, nil
}

// CreateActivity creates a new study activity
func (s *StudyActivityService) CreateActivity(activity *models.StudyActivity) error {
	activity.CreatedAt = time.Now()

	result, err := s.db.Exec(`
		INSERT INTO study_activities (group_id, activity_type, created_at)
		VALUES (?, ?, ?)`,
		activity.GroupID,
		activity.ActivityType,
		activity.CreatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	activity.ID = id
	return nil
}

// UpdateActivity updates an existing study activity
func (s *StudyActivityService) UpdateActivity(activity *models.StudyActivity) error {
	result, err := s.db.Exec(`
		UPDATE study_activities 
		SET group_id = ?, activity_type = ?
		WHERE id = ?`,
		activity.GroupID,
		activity.ActivityType,
		activity.ID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("study activity not found")
	}

	return nil
}

// DeleteActivity deletes a study activity by ID
func (s *StudyActivityService) DeleteActivity(id int64) error {
	result, err := s.db.Exec("DELETE FROM study_activities WHERE id = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("study activity not found")
	}

	return nil
}

// GetActivitySessions retrieves all study sessions for an activity
func (s *StudyActivityService) GetActivitySessions(activityID int64) ([]models.StudySession, error) {
	rows, err := s.db.Query(`
		SELECT id, start_time, end_time, score, status, study_activity_id
		FROM study_sessions
		WHERE study_activity_id = ?
		ORDER BY start_time DESC`,
		activityID,
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
