package service

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/erans/lang-portal/internal/models"
)

// WordService handles business logic for words
type WordService struct {
	db *sql.DB
}

// NewWordService creates a new WordService
func NewWordService(db *sql.DB) *WordService {
	return &WordService{db: db}
}

// GetWord retrieves a word by ID
func (s *WordService) GetWord(id int64) (*models.Word, error) {
	var word models.Word
	var partsJSON string

	err := s.db.QueryRow(
		"SELECT id, japanese, romaji, english, parts FROM words WHERE id = ?",
		id,
	).Scan(&word.ID, &word.Japanese, &word.Romaji, &word.English, &partsJSON)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("word not found")
		}
		return nil, err
	}

	// Parse JSON parts
	if err := json.Unmarshal([]byte(partsJSON), &word.Parts); err != nil {
		return nil, err
	}

	return &word, nil
}

// ListWords retrieves a paginated list of words
func (s *WordService) ListWords(offset, limit int) (*models.ListResult, error) {
	// Get total count first
	var totalItems int64
	err := s.db.QueryRow("SELECT COUNT(*) FROM words").Scan(&totalItems)
	if err != nil {
		return nil, err
	}

	// Get paginated words
	rows, err := s.db.Query(
		"SELECT id, japanese, romaji, english, parts FROM words LIMIT ? OFFSET ?",
		limit, offset,
	)
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

		if err := json.Unmarshal([]byte(partsJSON), &word.Parts); err != nil {
			return nil, err
		}

		words = append(words, word)
	}

	return &models.ListResult{
		Items:      words,
		TotalItems: totalItems,
	}, nil
}

// CreateWord creates a new word
func (s *WordService) CreateWord(word *models.Word) error {
	partsJSON, err := json.Marshal(word.Parts)
	if err != nil {
		return err
	}

	result, err := s.db.Exec(
		"INSERT INTO words (japanese, romaji, english, parts) VALUES (?, ?, ?, ?)",
		word.Japanese, word.Romaji, word.English, string(partsJSON),
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	word.ID = id
	return nil
}

// UpdateWord updates an existing word
func (s *WordService) UpdateWord(word *models.Word) error {
	partsJSON, err := json.Marshal(word.Parts)
	if err != nil {
		return err
	}

	result, err := s.db.Exec(
		"UPDATE words SET japanese = ?, romaji = ?, english = ?, parts = ? WHERE id = ?",
		word.Japanese, word.Romaji, word.English, string(partsJSON), word.ID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("word not found")
	}

	return nil
}

// DeleteWord deletes a word by ID
func (s *WordService) DeleteWord(id int64) error {
	result, err := s.db.Exec("DELETE FROM words WHERE id = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("word not found")
	}

	return nil
}
