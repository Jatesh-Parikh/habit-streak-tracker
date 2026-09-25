package resolvers

import (
	"context"
	"habit-streak-tracker/internal/models"
	"time"
)

func (r *userResolver) CreatedAt(ctx context.Context, obj *models.User) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *userResolver) UpdatedAt(ctx context.Context, obj *models.User) (string, error) {
	return obj.UpdatedAt.Format(time.RFC3339), nil
}
