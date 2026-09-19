package repository

import (
	"database/sql"
	"habit-streak-tracker/internal/models"

	// "errors"
	"fmt"
	"strings"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(name string, email string, hashedPassword string) (*models.User, error) {
	now := time.Now()

	var id string

	err := r.DB.QueryRow("SELECT lower(hex(randomblob(16)))").Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("Failed to generate user ID: %w", err)
	}

	_, err = r.DB.Exec("INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", id, name, email, hashedPassword, now, now)

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil, fmt.Errorf("Email already registered")
		}

		return nil, fmt.Errorf("Failed to create user: %w", err)
	}

	var user models.User

	err = r.DB.QueryRow("SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = ?", id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve created user: %w", err)
	}

	return &user, nil
}
