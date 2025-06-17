package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {

	//"3456,Ходьба,3h00m" - пример корректных входных данных

	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("incorrect number of parameters.\n expected 3 values.\n actual %d values", len(parts))
	}

	stepsCount, err := strconv.Atoi(parts[0])

	if stepsCount <= 0 {
		return 0, "", 0, errors.New("error: wrong value for steps. stepsCount <= 0")
	}

	if err != nil {
		return 0, "", 0, fmt.Errorf("error converting steps: %v", err)
	}

	tariningTime, err := time.ParseDuration(parts[2])

	if tariningTime <= 0 {
		return 0, "", 0, errors.New("error: wrong value for time. tariningTime <= 0")
	}
	if err != nil {
		return 0, "", 0, fmt.Errorf("error parsing time value: %v.", err)
	}

	traningType := parts[1]

	return stepsCount, traningType, tariningTime, err

}

func distance(steps int, height float64) float64 {

	dist := height * stepLengthCoefficient * float64(steps) / float64(mInKm)
	return dist
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {

	if duration <= 0 {
		return 0
	}

	dist := distance(steps, height)
	hours := duration.Hours()

	return (dist / hours)
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	var err error

	if steps > 0 && weight > 0 && height > 0 && duration > 0 {
		ms := meanSpeed(steps, height, duration)
		mins := duration.Minutes()
		return (weight * ms * mins / minInH), err
	}

	return 0, errors.New("incorrect number of parameters")
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	ccal, err := RunningSpentCalories(steps, weight, height, duration)

	return (ccal * walkingCaloriesCoefficient), err
}

func TrainingInfo(data string, weight, height float64) (string, error) {

	steps, activity, tariningTime, err := parseTraining(data)

	if err != nil {
		return "", err
	}

	dist := distance(steps, height)
	ms := meanSpeed(steps, height, tariningTime)

	var ccal float64

	switch {
	case activity == "Ходьба":
		ccal, err = WalkingSpentCalories(steps, weight, height, tariningTime)

	case activity == "Бег":
		ccal, err = RunningSpentCalories(steps, weight, height, tariningTime)

	default:
		err := errors.New("неизвестный тип тренировки") //unknown training type
		text := ""

		return text, err
	}

	if err != nil {
		return "", err
	}

	text := fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n", activity, tariningTime.Minutes()/60, dist, ms, ccal)

	return text, err
}
