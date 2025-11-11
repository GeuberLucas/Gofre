package models

type ResetToken struct {
	ID        int64
	UserID    int64
	TokenHash string
	ExpiresAt int64
}