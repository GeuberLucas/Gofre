package service

import (
	"context"
	"time"

	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/models"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/repository"
)

// IEventTrackService define os métodos disponíveis para a lógica de rastreamento de eventos
type IEventTrackService interface {
	// MarkEventAsProcessed registra que um evento foi consumido com sucesso
	MarkEventAsProcessed(ctx context.Context, eventID string, aggregateID int, consumerGroup string) error

	// IsEventProcessed verifica se um evento já foi processado anteriormente
	IsEventProcessed(ctx context.Context, eventID string, consumer string) (bool, error)
}

type EventTrackService struct {
	repo repository.IEventTrackRepository
}

// NewEventTrackService cria uma nova instância do serviço injetando o repositório
func NewEventTrackService(repo repository.IEventTrackRepository) IEventTrackService {
	return &EventTrackService{
		repo: repo,
	}
}

// MarkEventAsProcessed cria o modelo e chama o repositório para salvar
func (s *EventTrackService) MarkEventAsProcessed(ctx context.Context, eventID string, aggregateID int, consumerGroup string) error {
	track := &models.EventTrack{
		EventID:       eventID,
		AggregateID:   aggregateID,
		ConsumerGroup: consumerGroup,
		ProcessedAt:   time.Now(), // Define o momento exato do processamento
	}

	return s.repo.Save(ctx, track)
}

// IsEventProcessed verifica a existência do evento
func (s *EventTrackService) IsEventProcessed(ctx context.Context, eventID string, consumer string) (bool, error) {
	return s.repo.Exists(ctx, eventID, consumer)
}
