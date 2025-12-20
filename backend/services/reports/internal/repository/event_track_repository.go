package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/models"
)

type IEventTrackRepository interface {
	Save(ctx context.Context, track *models.EventTrack) error
	Exists(ctx context.Context, eventID string) (bool, error)
}

type EventTrackRepository struct {
	db *sql.DB
}

func NewEventTrackRepository(conn *sql.DB) IEventTrackRepository {
	return &EventTrackRepository{db: conn}
}

func (r *EventTrackRepository) Save(ctx context.Context, track *models.EventTrack) error {
	query := `
        INSERT INTO public.event_track (event_id, aggregate_id, consumer_group, processed_at)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (event_id) DO NOTHING
    `

	_, err := r.db.ExecContext(
		ctx,
		query,
		track.EventID,
		track.AggregateID,
		track.ConsumerGroup,
		track.ProcessedAt,
	)
	if err != nil {
		return fmt.Errorf("erro ao salvar event_track: %w", err)
	}

	return nil
}

func (r *EventTrackRepository) Exists(ctx context.Context, eventID string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM public.event_track WHERE event_id = $1)`

	err := r.db.QueryRowContext(ctx, query, eventID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("erro ao verificar existência do evento: %w", err)
	}

	return exists, nil
}
