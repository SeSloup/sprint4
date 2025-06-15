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

	if len(parts) == 2 {
		stepsCount, err := strconv.Atoi(parts[0])

		if err != nil || stepsCount <= 0 {
			err = fmt.Errorf("error converting steps: %v. stepsCount = %d", err, stepsCount)
			return 0, 0, err
		}

		walkTime, err := time.ParseDuration(parts[1])
		if err != nil || walkTime <= 0 {
			err = fmt.Errorf("error parsing time value: %v. walkTime = %.2f", err, walkTime)
			return 0, 0, err
		}

		return stepsCount, walkTime, nil
	}

	err := fmt.Errorf("incorrect number of parameters. expected: \"3456,3h00m\"")
	return 0, 0, err

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
