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

	fmt.Println("-- Minute --")
	fmt.Println("Input: 1,2       →", validateMinute("1,2"))
	fmt.Println("Input: 1-8       →", validateMinute("1-8"))
	fmt.Println("Input: *         →", validateMinute("*"))
	fmt.Println("Input: 4/9       →", validateMinute("4/9"))

	fmt.Println("-- Hour --")
	fmt.Println("Input: 1,2,13    →", validateHour("1,2,13"))
	fmt.Println("Input: 0/2       →", validateHour("0/2"))
	fmt.Println("Input: 9-17      →", validateHour("9-17"))

	fmt.Println("-- Day of Month --")
	fmt.Println("Input: 1,2,13    →", validateDayOfMonth("1,2,13"))
	fmt.Println("Input: 1/2       →", validateDayOfMonth("1/2"))
	fmt.Println("Input: 9-17      →", validateDayOfMonth("9-17"))
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

func validateHour(hourField string) string {
	if strings.Contains(hourField, ",") {
		values := strings.Split(hourField, ",")
		str := ""

		for index := range values {
			values[index] = convert24To12(values[index])
		}

		for index, value := range values {
			str += value
			if index == len(values)-2 {
				str += " and " + values[index+1]
				break
			}
			str += ", "
		}
		return "At " + str
	} else if strings.Contains(hourField, "/") {
		stepStart, stepInterval := handleStep(hourField)
		return "every " + strconv.Itoa(stepInterval) + " hours, starting at " + convert24To12(strconv.Itoa(stepStart))
	} else if strings.Contains(hourField, "-") {
		rangeStart, rangeEnd := handleRange(hourField)
		return "between " + convert24To12(strconv.Itoa(rangeStart)) + " and " + convert24To12(strconv.Itoa(rangeEnd))
	}
	return "every " + hourField
}

func validateDayOfMonth(dayOfMonth string) string {
	if strings.Contains(dayOfMonth, ",") {
		values := strings.Split(dayOfMonth, ",")
		str := ""

		for index, value := range values {
			str += value
			if index == len(values)-2 {
				str += " and " + values[index+1]
				break
			}
			str += ", "
		}
		return "on day " + str + " of the month"
	} else if strings.Contains(dayOfMonth, "/") {
		stepStart, stepInterval := handleStep(dayOfMonth)
		if stepStart == 1 {
			return "every " + strconv.Itoa(stepInterval) + " days"
		}
		return "every " + strconv.Itoa(stepInterval) + " days, starting on day " + strconv.Itoa(stepStart) + " of the month"
	} else if strings.Contains(dayOfMonth, "-") {
		rangeStart, rangeEnd := handleRange(dayOfMonth)
		return "between day " + strconv.Itoa(rangeStart) + " and " + strconv.Itoa(rangeEnd) + " of the month"
	}
	return "on day " + dayOfMonth + " of the month"
}

func convert24To12(hr string) string {
	i, _ := strconv.Atoi(hr)
	if i < 12 {
		return strconv.Itoa(i) + ":00 AM"
	} else {
		i %= 12
		return strconv.Itoa(i) + ":00 PM"
	}
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
