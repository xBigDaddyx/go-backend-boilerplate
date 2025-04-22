package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Buyers interface {
		Create(context.Context, *Buyer) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Buyers: &BuyerStore{db},
		Users:  &UserStore{db},
	}
}
