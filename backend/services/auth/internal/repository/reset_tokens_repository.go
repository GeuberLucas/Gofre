package repository

import (
	"database/sql"

	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/models"
)

type ResetTokensRepository struct{
	db *sql.DB
}

func NewResetTokensRepository(db *sql.DB) *ResetTokensRepository {
	return &ResetTokensRepository{db: db}
}

func (r *ResetTokensRepository) CreateResetToken(token *models.ResetToken) error {
	var sqlCommand string="insert into auth.reset_tokens (user_id, hash_token, expires_at) values ($1,$2,$3)"
	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(token.UserID, token.TokenHash, token.ExpiresAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *ResetTokensRepository) GetResetTokenByTokenHash(tokenHash string) (models.ResetToken, error) {
	var resetToken models.ResetToken
	var sqlCommand string="select id, user_id, hash_token, expires_at from auth.reset_tokens where hash_token = $1"
	row := r.db.QueryRow(sqlCommand, tokenHash)
	err := row.Scan(&resetToken.ID, &resetToken.UserID, &resetToken.TokenHash, &resetToken.ExpiresAt)
	if err != nil {
		return resetToken, err
	}
	return resetToken, nil
}