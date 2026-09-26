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

func (r *queryResolver) Habits(ctx context.Context) ([]*models.Habit, error) {
	userID, ok := middleware.GetUserID(ctx)

	if !ok {
		return nil, fmt.Errorf("Unauthorized")
	}

	habits, err := r.HabitRepo.GetHabitsByUserID(userID)

	if err != nil {
		return nil, err
	}

	return habits, nil
}

func (r *queryResolver) Habit(ctx context.Context, id string) (*models.Habit, error) {
	userID, ok := middleware.GetUserID(ctx)

	if !ok {
		return nil, fmt.Errorf("Unauthorized")
	}

	habit, err := r.HabitRepo.GetHabitWithUserCheck(id, userID)

	if err != nil {
		return nil, err
	}

	return habit, nil
}

func (r *queryResolver) HabitLogs(ctx context.Context, habitID string) ([]*models.HabitLog, error) {
	userID, ok := middleware.GetUserID(ctx)

	if !ok {
		return nil, fmt.Errorf("Unauthorized")
	}

	_, err := r.HabitRepo.GetHabitWithUserCheck(habitID, userID)

	if err != nil {
		return nil, err
	}

	logs, err := r.HabitLogRepo.GetHabitLogsByHabitID(habitID)

	if err != nil {
		return nil, err
	}

	return logs, nil
}
