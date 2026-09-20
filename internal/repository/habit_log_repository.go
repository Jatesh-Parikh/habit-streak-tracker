package repository

import (
	// "habit-streak-tracker/internal/models"
	"database/sql"
	// "errors"
	"fmt"
	// "strings"
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
