package models

import "time"

type EventTrack struct {
	EventID       string    `json:"event_id"`
	AggregateID   int       `json:"aggregate_id"`
	ConsumerGroup string    `json:"consumer_group"`
	ProcessedAt   time.Time `json:"processed_at"`
}
