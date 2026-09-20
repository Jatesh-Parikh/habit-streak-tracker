package repository

import (
	"database/sql"
	"habit-streak-tracker/internal/models"

	// "errors"
	"fmt"
	"strings"
	"time"
)

type HabitLogRepository struct {
	DB *sql.DB
}

func NewHabitLogRepository(db *sql.DB) *HabitLogRepository {
	return &HabitLogRepository{DB: db}
}

func (r *HabitLogRepository) CheckDuplicateLog(habitID string, completedDate time.Time) (bool, error) {
	dateStr := completedDate.Format("2006-01-02")

	var count int

	err := r.DB.QueryRow("SELECT COUNT(*) FROM habit_logs WHERE habit_id = ? AND completed_date = ?", habitID, dateStr).Scan(&count)

	if err != nil {
		return false, fmt.Errorf("Failed to check duplicate log: %w", err)
	}

	return count > 0, nil
}

func (r *HabitLogRepository) CreateHabitLog(habitID string, completedDate time.Time) (*models.HabitLog, error) {
	now := time.Now()

	truncated := completedDate.UTC().Truncate(24 * time.Hour)
	dateStr := truncated.Format("2006-01-02")

	var id string

	err := r.DB.QueryRow("SELECT lower(hex(randomblob(16)))").Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("Failed to generate log ID: %w", err)
	}

	_, err = r.DB.Exec("INSERT INTO habit_logs (id, habit_id, completed_date, created_at) VALUES (?, ?, ?, ?)", id, habitID, dateStr, now)

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil, fmt.Errorf("You have already checked in for this habit today or on this date")
		}

		return nil, fmt.Errorf("Failed to create habit: %w", err)
	}

	var log models.HabitLog

	err = r.DB.QueryRow("SELECT id, habit_id, completed_date, created_at FROM habit_logs WHERE id = ?", id).Scan(
		&log.ID,
		&log.HabitID,
		&log.CompletedDate,
		&log.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve created log: %w", err)
	}

	return &log, nil
}
