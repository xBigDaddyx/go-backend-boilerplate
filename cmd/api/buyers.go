package main

import (
	"net/http"

	"github.com/xbigdaddyx/go-backend-boilerplate/internal/store"
)

type CreateBuyerPayload struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func (app *application) createBuyerHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateBuyerPayload
	if err := readJSON(w, r, &payload); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	buyer := &store.Buyer{
		Name:     payload.Name,
		IsActive: payload.IsActive,
	}

	ctx := r.Context()

	if err := app.store.Buyers.Create(ctx, buyer); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, buyer); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
