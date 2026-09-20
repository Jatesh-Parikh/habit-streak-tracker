package repository

import (
	"database/sql"
	"habit-streak-tracker/internal/models"

	"errors"
	"fmt"

	// "strings"
	"time"
)

type HabitRepository struct {
	DB *sql.DB
}

func NewHabitRepository(db *sql.DB) *HabitRepository {
	return &HabitRepository{DB: db}
}

func (r *HabitRepository) CreateHabit(userID string, name string, description string) (*models.Habit, error) {
	now := time.Now()

	var id string

	err := r.DB.QueryRow("SELECT lower(hex(randomblob(16)))").Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("Failed to generate habit ID: %w", err)
	}

	_, err = r.DB.Exec("INSERT INTO habits (id, user_id, name, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", id, userID, name, description, now, now)

	if err != nil {
		return nil, fmt.Errorf("Failed to create this habit: %w", err)
	}

	var habit models.Habit

	err = r.DB.QueryRow("SELECT id, user_id, name, description, created_at, updated_at FROM habits WHERE id = ?", id).Scan(
		&habit.ID,
		&habit.UserID,
		&habit.Name,
		&habit.Description,
		&habit.CreatedAt,
		&habit.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to retreive created habit: %w", err)
	}

	return &habit, nil
}

func (r *HabitRepository) GetHabitsByUserID(userID string) ([]*models.Habit, error) {
	rows, err := r.DB.Query("SELECT id, user_id, name, description, created_at, updated_at FROM habits WHERE user_id = ? ORDER BY created_at DESC", userID)

	if err != nil {
		return nil, fmt.Errorf("Failed to fetch habits: %w", err)
	}

	defer rows.Close()

	habits := make([]*models.Habit, 0)

	for rows.Next() {
		var habit models.Habit

		err := rows.Scan(
			&habit.ID,
			&habit.UserID,
			&habit.Name,
			&habit.Description,
			&habit.CreatedAt,
			&habit.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("Failed to scan habit: %w", err)
		}

		habits = append(habits, &habit)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating habits: %w", err)
	}

	return habits, nil
}

func (r *HabitRepository) GetHabitWithUserCheck(habitID string, userID string) (*models.Habit, error) {
	var habit models.Habit

	err := r.DB.QueryRow("SELECT id, user_id, name, description, created_at, updated_at FROM habits WHERE id = ? AND user_id = ?", habitID, userID).Scan(
		&habit.ID,
		&habit.UserID,
		&habit.Name,
		&habit.Description,
		&habit.CreatedAt,
		&habit.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Habit not found")
		}

		return nil, fmt.Errorf("Failed to fetch habit: %w", err)
	}

	return &habit, nil
}
