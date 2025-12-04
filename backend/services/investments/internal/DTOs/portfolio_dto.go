package dtos

import (
	"time"
)

type Portfolio struct {
	Id           uint
	User_id      int
	Asset_id     uint
	Deposit_date time.Time
	Broker       string
	Amount       float64
	IsDone       bool
	Description  string
}
