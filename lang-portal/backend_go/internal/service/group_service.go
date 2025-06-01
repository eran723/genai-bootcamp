package service

import (
	"database/sql"
	"errors"

	"github.com/erans/lang-portal/internal/models"
)

// GroupService handles business logic for groups
type GroupService struct {
	db *sql.DB
}

// NewGroupService creates a new GroupService
func NewGroupService(db *sql.DB) *GroupService {
	return &GroupService{db: db}
}

// GetGroup retrieves a group by ID
func (s *GroupService) GetGroup(id int64) (*models.Group, error) {
	var group models.Group
	err := s.db.QueryRow(
		"SELECT id, name, description FROM groups WHERE id = ?",
		id,
	).Scan(&group.ID, &group.Name, &group.Description)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("group not found")
		}
		return nil, err
	}

	return &group, nil
}

// ListGroups retrieves a paginated list of groups
func (s *GroupService) ListGroups(offset, limit int) (*models.ListResult, error) {
	// Get total count first
	var totalItems int64
	err := s.db.QueryRow("SELECT COUNT(*) FROM groups").Scan(&totalItems)
	if err != nil {
		return nil, err
	}

	// Get paginated groups
	rows, err := s.db.Query(
		"SELECT id, name, description FROM groups LIMIT ? OFFSET ?",
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var group models.Group
		if err := rows.Scan(&group.ID, &group.Name, &group.Description); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	return &models.ListResult{
		Items:      groups,
		TotalItems: totalItems,
	}, nil
}

// CreateGroup creates a new group
func (s *GroupService) CreateGroup(group *models.Group) error {
	result, err := s.db.Exec(
		"INSERT INTO groups (name, description) VALUES (?, ?)",
		group.Name, group.Description,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	group.ID = id
	return nil
}

// UpdateGroup updates an existing group
func (s *GroupService) UpdateGroup(group *models.Group) error {
	result, err := s.db.Exec(
		"UPDATE groups SET name = ?, description = ? WHERE id = ?",
		group.Name, group.Description, group.ID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("group not found")
	}

	return nil
}

// DeleteGroup deletes a group by ID
func (s *GroupService) DeleteGroup(id int64) error {
	result, err := s.db.Exec("DELETE FROM groups WHERE id = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("group not found")
	}

	return nil
}

// GetGroupWords retrieves all words in a group
func (s *GroupService) GetGroupWords(groupID int64) ([]models.Word, error) {
	rows, err := s.db.Query(`
		SELECT w.id, w.japanese, w.romaji, w.english, w.parts
		FROM words w
		JOIN word_groups wg ON wg.word_id = w.id
		WHERE wg.group_id = ?
	`, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var words []models.Word
	for rows.Next() {
		var word models.Word
		var partsJSON string
		if err := rows.Scan(&word.ID, &word.Japanese, &word.Romaji, &word.English, &partsJSON); err != nil {
			return nil, err
		}
		words = append(words, word)
	}

	return words, nil
}

// GetGroupStudySessions retrieves study sessions for a group
func (s *GroupService) GetGroupStudySessions(groupID int64) ([]models.StudySession, error) {
	rows, err := s.db.Query(`
		SELECT ss.id, ss.start_time, ss.end_time, ss.score, ss.status, ss.study_activity_id
		FROM study_sessions ss
		JOIN study_activities sa ON sa.session_id = ss.id
		WHERE sa.group_id = ?
		ORDER BY ss.start_time DESC
	`, groupID)
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
