package models

import "time"

type EventTrack struct {
	EventID       string    `json:"event_id"`       //id event
	AggregateID   int       `json:"aggregate_id"`   // user owner
	ConsumerGroup string    `json:"consumer_group"` // consumer
	ProcessedAt   time.Time `json:"processed_at"`   // time processed
}
