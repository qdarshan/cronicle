package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	cronExpression := "5 12 1 4 4"
	// fmt.Println(cronExpression)

	if cronExpression == "" {
		err := errors.New("invalid expression")
		fmt.Println(err)
	}

	cronFields := strings.Split(cronExpression, " ")

	if len(cronFields) != 5 {
		err := errors.New("currently supports 5 standard cron fields")
		fmt.Println(err)
	}

	// minuteField := cronFields[0]
	// hourField := cronFields[1]
	// datOfMonthField := cronFields[2]
	// monthField := cronFields[3]
	// dayOfWeekField := cronFields[4]
	handleRange("JAN-DEC")
	handleRange("MON-WED")

	fmt.Println(validateMinute("1,2"))
	fmt.Println(validateMinute("1-8"))
	fmt.Println(validateMinute("*"))
	fmt.Println(validateMinute("4/9"))
}

func validateMinute(minuteField string) string {
	if strings.Contains(minuteField, "*") {
		return "every minute"
	} else if strings.Contains(minuteField, ",") {
		values := strings.Split(minuteField, ",")
		str := ""
		for i, value := range values {
			str += value
			if i == len(values)-2 {
				str += " and " + values[i+1]
				break
			}
			str += ", "
		}
		return "At " + str + " minutes past the hour"
	} else if strings.Contains(minuteField, "/") {
		stepStart, stepInterval := handleStep(minuteField)
		return "every " + strconv.Itoa(stepInterval) + " minutes, starting at " + strconv.Itoa(stepStart) + " minutes past the hour"
	} else if strings.Contains(minuteField, "-") {
		rangeStart, rangeEnd := handleRange(minuteField)
		return "Minutes " + strconv.Itoa(rangeStart) + " through " + strconv.Itoa(rangeEnd) + " past the hour"
	}
	return "every " + minuteField
}

func handleStep(field string) (int, int) {
	v := strings.Split(field, "/")
	return normalizeValue(v[0]), normalizeValue(v[1])
}

func handleRange(field string) (int, int) {
	v := strings.Split(field, "-")
	return normalizeValue(v[0]), normalizeValue(v[1])
}

func normalizeValue(value string) int {
	months := []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}
	weeks := []string{"SUN", "MON", "TUE", "WED", "THU", "FRI"}

	normalize, err := strconv.Atoi(value)
	if err != nil {
		for index, month := range months {
			if month == value {
				return index
			}
		}

		for index, week := range weeks {
			if week == value {
				return index
			}
		}
	}
	return normalize
}
