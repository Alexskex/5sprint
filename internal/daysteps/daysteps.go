package daysteps

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/errors"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dataSlice := strings.Split(datastring, ",")

	if len(dataSlice) != 2 {
		return errors.ErrInvalidArgumentCount
	}

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}

	if steps <= 0 {
		return errors.ErrZeroSteps
	}

	totalTime, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		return fmt.Errorf("invalid time format: %w", err)
	}
	if totalTime <= 0 {
		return errors.ErrZeroDuration
	}

	ds.Steps = steps
	ds.Duration = totalTime

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", errors.ErrZeroSteps
	}

	pathKilometers := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
result := fmt.Sprintf(
    "Количество шагов: %v.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
    ds.Steps,
    pathKilometers,
    calories,
)
return result, nil
}
