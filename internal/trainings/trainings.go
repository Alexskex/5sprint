package trainings

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/errors"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type Training struct {
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) (err error) {
	dataSlice := strings.Split(datastring, ",")

	if len(dataSlice) != 3 {
		return errors.ErrInvalidArgumentCount
	}

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}
	if steps <= 0 {
		return errors.ErrZeroSteps
	}

	totalTime, err := time.ParseDuration(dataSlice[2])
	if err != nil {
		return fmt.Errorf("invalid time format: %w", err)
	}
	if totalTime <= 0 {
		return errors.ErrZeroDuration
	}

	t.Steps = steps
	t.TrainingType = dataSlice[1]
	t.Duration = totalTime

	return nil
}

func (t Training) ActionInfo() (string, error) {
	var totalCalories float64

	switch t.TrainingType {
	case "Ходьба":
		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		} else {
			totalCalories = calories
		}
	case "Бег":
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		} else {
			totalCalories = calories
		}
	default:
		return "", errors.ErrUnknownTraining
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	speedAvg := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	

	result := fmt.Sprintf(
    "Тип тренировки: %s\n"+
    "Длительность: %.2f ч.\n"+
    "Дистанция: %.2f км.\n"+
    "Скорость: %.2f км/ч\n"+
    "Сожгли калорий: %.2f\n",
    t.TrainingType,
    t.Duration.Hours(),
    distance,
    speedAvg,
    totalCalories,
)

	return result, nil
}

