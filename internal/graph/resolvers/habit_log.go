package resolvers

import (
	"context"
	"habit-streak-tracker/internal/models"

	"fmt"
	"time"
)

func (r *habitLogResolver) CompletedDate(ctx context.Context, obj *models.HabitLog) (string, error) {
	return obj.CompletedDate.Format("2006-01-02"), nil
}

func (r *habitLogResolver) CreatedAt(ctx context.Context, obj *models.HabitLog) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *habitLogResolver) Habit(ctx context.Context, obj *models.HabitLog) (*models.Habit, error) {
	habit, err := r.HabitRepo.GetHabitByID(obj.HabitID)

	if err != nil {
		return nil, err
	}

	if habit == nil {
		return nil, fmt.Errorf("Habit not found")
	}

	return habit, nil
}
