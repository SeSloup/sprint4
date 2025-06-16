package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {

	parts := strings.Split(data, ",")

	if len(parts) != 2 {
		err := fmt.Errorf("incorrect number of parameters.\n expected 3 values.\n actual %d values", len(parts))
		return 0, 0, err
	}
	stepsCount, err := strconv.Atoi(parts[0])

	if stepsCount <= 0 {
		err = fmt.Errorf("error: wrong value for steps. stepsCount = %d", stepsCount)
		return 0, 0, err
	}

	if err != nil {
		err = fmt.Errorf("error converting steps: %v", err)
		return 0, 0, err
	}

	walkTime, err := time.ParseDuration(parts[1])
	if walkTime <= 0 {
		err = fmt.Errorf("error: wrong value for time. walkTime = %.2f", walkTime)
		return 0, 0, err
	}
	if err != nil {
		err = fmt.Errorf("error parsing time value: %v.", err)
		return 0, 0, err
	}

	return stepsCount, walkTime, nil

}

func DayActionInfo(data string, weight, height float64) string {

	// "1000,1h30m" - пример корректных входных данных
	steps, walkTime, _ := parsePackage(data)
	dist := float64(steps) * stepLength / mInKm

	ccal, _ := spentcalories.WalkingSpentCalories(steps, weight, height, walkTime)

	return fmt.Sprintf("Количество шагов: %d./n"+
		"Дистанция составила %.2f км./n"+
		"Вы сожгли %.2f ккал./n", steps, dist, ccal)

}
