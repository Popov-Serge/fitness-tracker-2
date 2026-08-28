package trainings

import (
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"

	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return errors.New("invalid data")
	}
	t.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}
	if t.Steps <= 0 {
		return errors.New("Количество шагов должно быть больше 0")
	}

	t.Duration, err = time.ParseDuration(data[2])
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}
	if t.Duration <= 0 {
		return errors.New("Длительность тренировки должна быть больше 0")
	}

	t.TrainingType = data[1]

	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.TrainingType != "Бег" && t.TrainingType != "Ходьба" {
		return "", errors.New("неизвестный тип тренировки")
	}

	trainingType := fmt.Sprintf("Тип тренировки: %s\n", t.TrainingType)
	duration := fmt.Sprintf("Длительность: %.2f ч.\n", t.Duration.Hours())
	distance := fmt.Sprintf("Дистанция: %.2f км.\n", spentenergy.Distance(t.Steps, t.Height))
	speed := fmt.Sprintf("Скорость: %.2f км/ч\n", spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration))
	var burntCaloriesStr string
	var burntCalories float64

	var err error
	if t.TrainingType == "Бег" {
		burntCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		burntCaloriesStr = fmt.Sprintf("Сожгли калорий: %.2f\n", burntCalories)
	} else {
		burntCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		burntCaloriesStr = fmt.Sprintf("Сожгли калорий: %.2f\n", burntCalories)
	}

	if err != nil {
		return "", fmt.Errorf("Calories error: %w", err)
	}

	var builder strings.Builder
	builder.WriteString(trainingType)
	builder.WriteString(duration)
	builder.WriteString(distance)
	builder.WriteString(speed)
	builder.WriteString(burntCaloriesStr)

	return builder.String(), nil
}
