package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше 0")
	}

	if weight <= 0 {
		return 0, errors.New("вес должен быть больше 0")
	}

	if height <= 0 {
		return 0, errors.New("рост должен быть больше 0")
	}

	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше 0")
	}

	durationInMinutes := duration.Minutes()
	meanSpeed := MeanSpeed(steps, height, duration)

	return (weight * meanSpeed * durationInMinutes * walkingCaloriesCoefficient) / minInH, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше 0")
	}

	if weight <= 0 {
		return 0, errors.New("вес должен быть больше 0")
	}

	if height <= 0 {
		return 0, errors.New("рост должен быть больше 0")
	}

	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше 0")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration < 0 {
		return 0
	}
	distance := Distance(steps, height)
	durationHours := duration.Hours()
	if durationHours <= 0 {
		return 0
	}

	return distance / durationHours
}

func Distance(steps int, height float64) float64 {
	stepLength := stepLengthCoefficient * height
	distance := stepLength * float64(steps)
	return distance / mInKm
}
