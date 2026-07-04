package spentenergy

import (
	"time"

	"github.com/Yandex-Practicum/tracker/internal/errors"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	cal, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	return cal * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.ErrZeroSteps
	}
	if weight <= 0 {
		return 0, errors.ErrZeroWeight
	}
	if height <= 0 {
		return 0, errors.ErrZeroHeigth
	}
	if duration <= 0 {
		return 0, errors.ErrZeroDuration
	}

	speedAvg := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := (weight * speedAvg * durationInMinutes) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	totalDistance := Distance(steps, height)
	return totalDistance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := stepLengthCoefficient * height
	pathLength := stepLength * float64(steps)
	return pathLength / mInKm
}


