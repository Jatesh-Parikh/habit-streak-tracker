package resolvers

import (
	"habit-streak-tracker/internal/models"
	// "habit-streak-tracker/internal/utils"
	"context"
	// "fmt"
	"time"
)

func (r *habitResolver) CreatedAt(ctx context.Context, obj *models.Habit) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *habitResolver) UpdatedAt(ctx context.Context, obj *models.Habit) (string, error) {
	return obj.UpdatedAt.Format(time.RFC3339), nil
}
