package repository

import (
	"database/sql"
	"habit-streak-tracker/internal/models"

	"errors"
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

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?", email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Invalid email or password")
		}

		return nil, fmt.Errorf("Failed to find user: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User

	err := r.DB.QueryRow("SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?", id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("User not found")
		}

		return nil, fmt.Errorf("Failed to find user: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(id string, name *string, email *string, password *string) (*models.User, error) {
	var setClauses []string

	var args []interface{}

	if name != nil {
		setClauses = append(setClauses, "name = ?")
		args = append(args, *name)
	}

	if email != nil {
		setClauses = append(setClauses, "email = ?")
		args = append(args, *email)
	}

	if password != nil {
		setClauses = append(setClauses, "password = ?")
		args = append(args, *password)
	}

	if len(setClauses) == 0 {
		return nil, fmt.Errorf("No fields provided to update")
	}

	now := time.Now()
	setClauses = append(setClauses, "updated_at = ?")
	args = append(args, now)

	setClause := strings.Join(setClauses, ", ")

	query := fmt.Sprintf("UPDATE users SET %s WHERE id = ?", setClause)
	args = append(args, id)

	_, err := r.DB.Exec(query, args...)

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil, fmt.Errorf("Email already in use")
		}

		return nil, fmt.Errorf("Failed to update user: %w", err)
	}

	return r.GetUserByID(id)
}

func (r *UserRepository) DeleteUser(id string) (bool, error) {
	result, err := r.DB.Exec("DELETE FROM users WHERE id = ?", id)

	if err != nil {
		return false, fmt.Errorf("Failed to delete the user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return false, fmt.Errorf("Failed to confirm deletion: %w", err)
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
