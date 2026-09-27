package resolvers

import (
	"context"
	"fmt"
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
