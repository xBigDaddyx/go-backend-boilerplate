package store

import (
	"context"
	"database/sql"
)

type Buyer struct {
	ID        int64  `json:"id"`
	Logo      string `json:"logo"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type BuyerStore struct {
	db *sql.DB
}

func (s *BuyerStore) Create(ctx context.Context, buyer *Buyer) error {
	query := `
	INSERT INTO t_organization_buyers (name,is_active)
	VALUES ($1,$2) RETURNING id, created_at,updated_at
	`
	err := s.db.QueryRowContext(
		ctx,
		query,
		buyer.Name,
		buyer.IsActive,
	).Scan(
		&buyer.ID,
		&buyer.CreatedAt,
		&buyer.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
