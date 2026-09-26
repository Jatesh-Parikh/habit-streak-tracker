package resolvers

import (
	"context"
	"fmt"
	"habit-streak-tracker/internal/middleware"
	"habit-streak-tracker/internal/models"
)

func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
	userID, ok := middleware.GetUserID(ctx)

	if !ok {
		return nil, fmt.Errorf("Unauthorized")
	}

	user, err := r.UserRepo.GetUserByID(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}
