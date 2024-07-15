package services

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// NextDate() provides with new date for task
func NextDate(now time.Time, date string, repeat string) (string, error) {
	const op = "services.NextDate"

	dateTime, err := time.Parse("20060102", date)
	if err != nil {
		return "", fmt.Errorf("Некорректный формат даты. %s:%v", op, err)
	}
	if len(repeat) == 0 {
		return "", fmt.Errorf("Не задана периодичность. %s:%v", op, err)
	}
	return condition(now, dateTime, repeat)
}

// condition() checks condition and provides new date for task depends on condition
func condition(now time.Time, dateTime time.Time, repeat string) (string, error) {
	const op = "services.Repeat"
	repeatSlice := strings.Split(repeat, " ")
	switch repeatSlice[0] {
	case "y":
		return conditionForYear(repeatSlice, now, dateTime)
	case "d":
		return conditionForDay(repeatSlice, now, dateTime)

	default:
		return "", fmt.Errorf("Некорректный формат параметра repeat. %s:%s", op, repeatSlice)
	}
}

// conditionForYear() calculates date for task in next year
func conditionForYear(repeatSlice []string, now time.Time, dateTime time.Time) (string, error) {
	const op = "services.conditionForYear"

	if len(repeatSlice) != 1 {
		return "", fmt.Errorf("Некорректный формат параметра repeat. %s:%s", op, repeatSlice)
	}
	dateTime = dateTime.AddDate(1, 0, 0)
	for dateTime.Format("20060102") <= now.Format("20060102") {
		dateTime = dateTime.AddDate(1, 0, 0)
	}
	return dateTime.Format("20060102"), nil
}

// conditionForYear() calculates date for task in next days
func conditionForDay(repeatSlice []string, now time.Time, dateTime time.Time) (string, error) {
	const op = "services.conditionForDay"

	if len(repeatSlice) != 2 {
		return "", fmt.Errorf("Некорректный формат параметра repeat. %s:%s", op, repeatSlice)
	}

	duration, err := strconv.Atoi(repeatSlice[1])
	if err != nil {
		return "", fmt.Errorf("Некорректный формат параметра repeat. %s:%s", op, repeatSlice)
	}
	if duration <= 0 {
		return "", fmt.Errorf("Длительность в днях не может быть меньше ноля. %s:%s", op, duration)
	}
	if duration > 400 {
		return "", fmt.Errorf("Максимально допустимая длительность в днях составляет 400. %s:%s", op, duration)
	}

	dateTime = dateTime.AddDate(0, 0, duration)
	for dateTime.Format("20060102") <= now.Format("20060102") {
		dateTime = dateTime.AddDate(0, 0, duration)
	}
	return dateTime.Format("20060102"), nil
}

//TODO добавить обработчики для недели (w) и месяца (m)

/* case "w":
	if len(repeatSlice) != 2 {
		return "", fmt.Errorf("invalid format of repeat's string")
	}
	daysMap := make(map[string]struct{})
	daysOfWeek := strings.Split(repeatSlice[1], ",")
	for _, day := range daysOfWeek {
		if day > "7" || day < "1" {
			return "", fmt.Errorf("invalid format of week's parametr")
		}
		daysMap[day] = struct{}{}
	}
	var checkDate time.Time
	if now.Unix() > d.Unix() {
		checkDate = now
	} else {
		checkDate = d
	}
	for {
		if _, ok := daysMap[checkDate.Weekday().String()]; ok {
			return checkDate.Format("20060102"), nil
		}
		checkDate = checkDate.AddDate(0, 0, 1)
	}
case "m":
	//TODO m cycle
	if len(repeatSlice) == 2 {
		params := strings.Split(repeatSlice[1], ",")
		for _, param := range params {
			if numParam, err := strconv.Atoi(param); err != nil || numParam < (-31) || numParam > 31 {
				return "", fmt.Errorf("invalid format of month's parametr")
			}

		}
	}
	if len(repeatSlice) == 3 {

	}
	return "", fmt.Errorf("invalid format of repeat's string")*/
