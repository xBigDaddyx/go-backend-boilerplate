package store

import (
	"context"
	"database/sql"
)

type SewingLine struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Type        string `json:"name"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type SewingLineStore struct {
	db *sql.DB
}

func (s *SewingLineStore) Create(ctx context.Context, line *SewingLine) error {
	query := `
	INSERT INTO sewing_lines (id,display_name,type,is_active)
	VALUE ($1,$2,$3,$4) RETURNING id, created_at,updated_at
	`
	err := s.db.QueryRowContext(
		ctx,
		query,
		line.ID,
		line.DisplayName,
		line.Type,
		line.IsActive,
	).Scan(
		&line.CreatedAt,
		&line.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
