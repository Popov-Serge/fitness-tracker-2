package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dayData := strings.Split(datastring, ",")
	if len(dayData) != 2 {
		return fmt.Errorf("Некорректные данные")
	}

	ds.Steps, err = strconv.Atoi(dayData[0])
	if err != nil {
		return err
	}

	if ds.Steps <= 0 {
		return errors.New("количество шагов должно быть больше 0")
	}

	ds.Duration, err = time.ParseDuration(dayData[1])
	if err != nil {
		return err
	}

	if ds.Duration <= 0 {
		return errors.New("продолжительность должна быть больше 0")
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	var builder strings.Builder

	stepsCount := fmt.Sprintf("Количество шагов: %d.\n", ds.Steps)
	builder.WriteString(stepsCount)
	distance := fmt.Sprintf("Дистанция составила %.2f км.\n", spentenergy.Distance(ds.Steps, ds.Height))
	builder.WriteString(distance)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	caloriesStr := fmt.Sprintf("Вы сожгли %.2f ккал.\n", calories)
	builder.WriteString(caloriesStr)

	return builder.String(), nil
}
