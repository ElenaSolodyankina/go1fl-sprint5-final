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
	if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {
		return 0.0, errors.New("invalid parameters")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	spentCalories := (weight * meanSpeed * duration.Minutes()) / minInH * walkingCaloriesCoefficient

	return spentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {
		return 0.0, errors.New("invalid parameters")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	spentCalories := (weight * meanSpeed * duration.Minutes()) / minInH

	return spentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	meanSpeed := distance / duration.Hours()

	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	distance := (height * stepLengthCoefficient * float64(steps)) / float64(mInKm)

	return distance
}
