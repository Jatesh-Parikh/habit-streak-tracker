package resolvers

import (
	"context"
	"habit-streak-tracker/internal/models"

	// "fmt"
	"time"
)

func (r *habitLogResolver) CompletedDate(ctx context.Context, obj *models.HabitLog) (string, error) {
	return obj.CompletedDate.Format("2006-01-02"), nil
}

func (r *habitLogResolver) CreatedAt(ctx context.Context, obj *models.HabitLog) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}
