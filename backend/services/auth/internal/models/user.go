package models

import (
	"time"
)

type User struct {
	ID        int64
	Username  string
	Name      string
	LastName  string
	Cellphone string
	Email     string
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}