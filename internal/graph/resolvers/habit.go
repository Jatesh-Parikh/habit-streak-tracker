package resolvers

import (
	"habit-streak-tracker/internal/models"
	// "habit-streak-tracker/internal/utils"
	"context"
	"fmt"
	"time"
)

func (r *habitResolver) CreatedAt(ctx context.Context, obj *models.Habit) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *habitResolver) UpdatedAt(ctx context.Context, obj *models.Habit) (string, error) {
	return obj.UpdatedAt.Format(time.RFC3339), nil
}

func (r *habitResolver) User(ctx context.Context, obj *models.Habit) (*models.User, error) {
	user, err := r.UserRepo.GetUserByID(obj.UserID)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("User not found")
	}

	return user, nil
}

func (r *habitResolver) Logs(ctx context.Context, obj *models.Habit) ([]*models.HabitLog, error) {
	logs, err := r.HabitLogRepo.GetHabitLogsByHabitID(obj.ID)

	if err != nil {
		return nil, err
	}

	return logs, nil
}
