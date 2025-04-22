package store

import (
	"context"
	"database/sql"
)

type SewingPlan struct {
	ID           int64  `json:"id"`
	BuyerId      int64  `json:"buyer_id"`
	SewingLineId string `json:"sewing_line_id"`
	Contract     string `json:"contract"`
	Style        string `json:"style"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type SewingPlanStore struct {
	db *sql.DB
}

func (s *SewingPlanStore) Create(ctx context.Context, plan *SewingPlan) error {
	query := `
	INSERT INTO sewing_plans (buyer_id,sewing_line_id,contract,style,start_date,end_date,status)
	VALUE ($1,$2,$3,$4,$5,$6,$7) RETURNING id, created_at,updated_at
	`
	err := s.db.QueryRowContext(
		ctx,
		query,
		plan.BuyerId,
		plan.SewingLineId,
		plan.Contract,
		plan.Style,
		plan.StartDate,
		plan.EndDate,
		plan.Status,
	).Scan(
		&plan.ID,
		&plan.CreatedAt,
		&plan.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
