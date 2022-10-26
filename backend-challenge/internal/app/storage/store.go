package storage

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

const (
	databaseURL = "host=localhost dbname=challenge_test sslmode=disable"
)

// New ...
func NewStorage() (*Storage, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) FindByEmail(email string) (*User, error) {
	u := &User{}
	if err := s.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return u, nil
}

func (s *Storage) SessionSave(userId string) (string, error) {
	sessionId, err := encryptString(userId)
	if err != nil {
		return "", err
	}

	if err := s.db.QueryRow(
		"UPDATE users SET sessionId = $1 where userId = $2",
		sessionId, userId,
	).Scan(); err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", err
	}
	return sessionId, nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil

}
