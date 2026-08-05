package security

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type ResetToken struct {
	TokenHash string
	ExpiresAt time.Time
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPasswordHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

func CreateResetToken(sizeHash int) (string, ResetToken, error) {
	b := make([]byte, sizeHash)
	_, err := rand.Read(b)
	if err != nil {
		return "",ResetToken{} ,err
	}
	token := base64.URLEncoding.EncodeToString(b)

	tokenHash := HashToken(token)

	return token, ResetToken{
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}, nil
}

func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(hash[:])
}