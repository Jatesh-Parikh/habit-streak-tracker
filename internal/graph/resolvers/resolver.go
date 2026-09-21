package resolvers

import "habit-streak-tracker/internal/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	UserRepo     *repository.UserRepository
	HabitRepo    *repository.HabitRepository
	HabitLogRepo *repository.HabitLogRepository
}
