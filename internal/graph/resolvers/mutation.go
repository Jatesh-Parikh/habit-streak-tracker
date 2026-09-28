package resolvers

import (
	"context"
	"fmt"
	"habit-streak-tracker/internal/middleware"
	"habit-streak-tracker/internal/models"
	"habit-streak-tracker/internal/utils"
	"os"
)

func (r *mutationResolver) Register(ctx context.Context, name string, email string, password string) (*models.AuthPayload, error) {
	err := utils.ValidateName(name)

	if err != nil {
		return nil, fmt.Errorf("Invalid name :%w", err)
	}

	err = utils.ValidateEmail(email)

	if err != nil {
		return nil, fmt.Errorf("Invalid email :%w", err)
	}

	err = utils.ValidatePasswordStrength(password)

	if err != nil {
		return nil, fmt.Errorf("Invalid password :%w", err)
	}

	hashedPassword, err := utils.HashPassword(password)

	if err != nil {
		return nil, fmt.Errorf("Failed to hash password: %w", err)
	}

	user, err := r.UserRepo.CreateUser(name, email, hashedPassword)

	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateJWT(user.ID, os.Getenv("JWT_SECRET"))

	if err != nil {
		return nil, fmt.Errorf("Failed to generate token: %w", err)
	}

	return &models.AuthPayload{
		Token: token,
		User:  user,
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*models.AuthPayload, error) {
	user, err := r.UserRepo.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	err = utils.ComparePassword(user.Password, password)

	if err != nil {
		return nil, fmt.Errorf("Invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID, os.Getenv("JWT_SECRET"))

	if err != nil {
		return nil, fmt.Errorf("Failed to generate token: %w", err)
	}

	return &models.AuthPayload{
		Token: token,
		User:  user,
	}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, name *string, email *string, password *string) (*models.User, error) {
	userID, ok := middleware.GetUserID(ctx)

	if !ok {
		return nil, fmt.Errorf("Unauthorized")
	}

	if name != nil {
		if err := utils.ValidateName(*name); err != nil {
			return nil, fmt.Errorf("Invalid name: %w", err)
		}
	}

	if email != nil {
		if err := utils.ValidateEmail(*email); err != nil {
			return nil, fmt.Errorf("Invalid email: %w", err)
		}
	}

	var hashedPassword *string

	if password != nil {
		if err := utils.ValidatePasswordStrength(*password); err != nil {
			return nil, fmt.Errorf("Invalid password: %w", err)
		}

		hash, err := utils.HashPassword(*password)

		if err != nil {
			return nil, fmt.Errorf("Failed to hash password: %w", err)
		}

		hashedPassword = &hash
	}

	user, err := r.UserRepo.UpdateUser(userID, name, email, hashedPassword)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context) (bool, error) {
	userID, ok := middleware.GetUserID(ctx)

	if !ok {
		return false, fmt.Errorf("Unauthorized")
	}

	deleted, err := r.UserRepo.DeleteUser(userID)

	if err != nil {
		return false, err
	}

	if !deleted {
		return false, fmt.Errorf("User not found")
	}

	return true, nil
}

func (r *mutationResolver) CreateHabit(ctx context.Context, name string, description string) (*models.Habit, error) {
	userID, ok := middleware.GetUserID(ctx)

	if !ok {
		return nil, fmt.Errorf("Unauthorized")
	}

	if err := utils.ValidateName(name); err != nil {
		return nil, fmt.Errorf("Invalid habit name: %w", err)
	}

	if err := utils.ValidateDescription(description); err != nil {
		return nil, fmt.Errorf("Invalid description: %w", err)
	}

	habit, err := r.HabitRepo.CreateHabit(userID, name, description)

	if err != nil {
		return nil, err
	}

	return habit, nil
}

func (r *mutationResolver) UpdateHabit(ctx context.Context, id string, name *string, description *string) (*models.Habit, error) {
	userID, ok := middleware.GetUserID(ctx)

	if !ok {
		return nil, fmt.Errorf("Unauthorized")
	}

	if name != nil {
		if err := utils.ValidateName(*name); err != nil {
			return nil, fmt.Errorf("Invalid name: %w", err)
		}
	}

	if description != nil {
		if err := utils.ValidateDescription(*description); err != nil {
			return nil, fmt.Errorf("Invalid description: %w", err)
		}
	}

	habit, err := r.HabitRepo.UpdateHabit(id, userID, name, description)

	if err != nil {
		return nil, err
	}

	return habit, nil
}
