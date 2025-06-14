package daysteps

import (
	"fmt"
	"strconv"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	//"1000,1h30m"
	s := ""         //строка даты
	x := ""         //перебираемая буква
	stepsCount := 0 //количество шагов
	var walkTime time.Duration
	var err error

	for i := 0; i < len(data); i++ {
		x = string(data[i])
		if x != "," {
			s = s + x
			continue

		}
		if stepsCount == 0 {
			stepsCount, err = strconv.Atoi(s)
			if err != nil {
				return stepsCount, walkTime, err
			}
			s = ""
			continue

		}
	}
	/*1, Jan, January — месяц;
	  2 — число месяца;
	  3, 15 — час в 12- и 24-часовом формате соответственно;
	  4 — минуты;
	  5 — секунды;
	  06, 2006 — год;
	  -0700, Z0700, Z07:00, Z07 — часовой пояс;
	  Mon, Monday — день недели;
	  pm, PM — время суток;
	  MST — аббревиатура часового пояса.
	*/
	walkTime, err = time.ParseDuration(s)
	if err != nil {
		return stepsCount, walkTime, err
	}

	return stepsCount, walkTime, nil

}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	// "1000,1h30m"
	steps, walkTime, _ := parsePackage(data)
	dist := float64(steps) * stepLength / mInKm

	ccal, _ := WalkingSpentCalories(steps, weight, height, walkTime)

	return fmt.Sprintf("Количество шагов: %s. /n"+
		"Дистанция составила %d км. /n"+
		"Вы сожгли %d ккал. ", steps, dist, ccal)

}
