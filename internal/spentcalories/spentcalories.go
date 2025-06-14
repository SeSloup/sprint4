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
	// TODO: реализовать функцию
	//"3456,Ходьба,3h00m"
	stepsCount := 0 //количество шагов
	var traningType string
	var tariningTime time.Duration
	var err error

	parts := strings.Split(data, ",")
	if len(parts) == 3 {
		stepsCount, err = strconv.Atoi(parts[0])

		if err != nil || stepsCount <= 0 {
			err = fmt.Errorf("Ошибка конвертации шагов.")
			return 0, "", 0, err
		}

		traningType = parts[1]

		tariningTime, err = time.ParseDuration(parts[2])
		if err != nil || tariningTime <= 0 {
			err = fmt.Errorf("Ошибка парсинга времени.")
			return 0, "", 0, err
		}

		return stepsCount, traningType, tariningTime, nil
	}

	err = fmt.Errorf("Количество параметров некорректно")
	return 0, "", 0, err

}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	var dist float64
	dist = height * stepLengthCoefficient * float64(steps) / float64(mInKm)
	return dist
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	dist := distance(steps, height)
	hours := duration.Hours()

	return (dist / hours)
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	var err error

	if steps > 0 && weight > 0 && height > 0 && duration > 0 {
		ms := meanSpeed(steps, height, duration)
		mins := duration.Minutes()
		return (weight * ms * mins / minInH), err
	}

	return 0, fmt.Errorf("Некорректные входные данные.")
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	ccal, err := RunningSpentCalories(steps, weight, height, duration)

	return (ccal * walkingCaloriesCoefficient), err
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	var err error
	var ccal float64
	var err2 error

	knownActivities := map[string]int{
		"Ходьба": 1,
		"Бег":    2,
	}

	steps, activity, tariningTime, err1 := parseTraining(data)
	dist := distance(steps, height)
	ms := meanSpeed(steps, height, tariningTime)

	if activity == "Ходьба" {
		ccal, err2 = WalkingSpentCalories(steps, weight, height, tariningTime)
	} else {
		ccal, err2 = RunningSpentCalories(steps, weight, height, tariningTime)
	}

	err = errors.Join(err1, err2)

	if _, ok := knownActivities[activity]; !ok {

		err3 := fmt.Errorf("неизвестный тип тренировки")
		text := ""

		err = errors.Join(err, err3)
		return text, err
	}

	text := fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n", activity, tariningTime.Minutes()/60, dist, ms, ccal)

	return text, err
}
