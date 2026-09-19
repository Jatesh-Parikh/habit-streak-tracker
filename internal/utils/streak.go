package utils

import (
	"habit-streak-tracker/internal/models"
	"time"
)

func CalculateCurrentStreak(logs []*models.HabitLog) int {
	if len(logs) == 0 {
		return 0
	}

	today := time.Now().UTC().Truncate(24 * time.Hour)
	yesterday := today.AddDate(0, 0, -1)

	firstLogDate := logs[0].CompletedDate

	if !firstLogDate.Equal(today) && !firstLogDate.Equal(yesterday) {
		return 0
	}

	streak := 1

	prevDate := firstLogDate

	for i := 1; i < len(logs); i++ {
		logDate := logs[i].CompletedDate
		expectedDate := prevDate.AddDate(0, 0, -1)

		if logDate.Equal(expectedDate) {
			streak++

			prevDate = logDate
		} else {
			break
		}
	}

	return streak
}

func CalculateLongestStreak(logs []*models.HabitLog) int {
	if len(logs) == 0 {
		return 0
	}

	for i, j := 0, len(logs)-1; i < j; i, j = i+1, j-1 {
		logs[i], logs[j] = logs[j], logs[i]
	}

	currentRun := 1
	longestRun := 1

	prevDate := logs[0].CompletedDate.UTC().Truncate(24 * time.Hour)

	for i := 1; i < len(logs); i++ {
		logDate := logs[i].CompletedDate.UTC().Truncate(24 * time.Hour)
		expectedDate := prevDate.AddDate(0, 0, 1)

		if logDate.Equal(expectedDate) {
			currentRun++

			if currentRun > longestRun {
				longestRun = currentRun
			}
		} else {
			currentRun = 1
		}

		prevDate = logDate
	}

	return longestRun
}
