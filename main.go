package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// unix corn
func main() {
	cronExpression := "5 12 1 4 4"
	// fmt.Println(cronExpression)

	if cronExpression == "" {
		log.Fatal("invalid expression")
	}

	cronFields := strings.Split(cronExpression, " ")
	if len(cronFields) != 5 {
		log.Fatal("currently supports 5 standard cron fields")
	}

	fmt.Println("-- Minute --")
	fmt.Println("Input: 1         →", validateMinute("1"))
	fmt.Println("Input: 1,2       →", validateMinute("1,2"))
	fmt.Println("Input: 1-8       →", validateMinute("1-8"))
	fmt.Println("Input: *         →", validateMinute("*"))
	fmt.Println("Input: 4/9       →", validateMinute("4/9"))

	fmt.Println("-- Hour --")
	fmt.Println("Input: 1         →", validateHour("1"))
	fmt.Println("Input: 1,2,13    →", validateHour("1,2,13"))
	fmt.Println("Input: 0/2       →", validateHour("0/2"))
	fmt.Println("Input: 9-17      →", validateHour("9-17"))

	fmt.Println("-- Day of Month --")
	fmt.Println("Input: 1         →", validateDayOfMonth("1"))
	fmt.Println("Input: 1,2,13    →", validateDayOfMonth("1,2,13"))
	fmt.Println("Input: 1/2       →", validateDayOfMonth("1/2"))
	fmt.Println("Input: 3/5       →", validateDayOfMonth("3/5"))
	fmt.Println("Input: 9-17      →", validateDayOfMonth("9-17"))
	// fmt.Println("Input: L         →", validateDayOfMonth("L"))

	fmt.Println("-- Month --")
	fmt.Println("Input: 1         →", validateMonth("1"))
	fmt.Println("Input: 1,2,11    →", validateMonth("1,2,11"))
	fmt.Println("Input: 1/3       →", validateMonth("1/3"))
	fmt.Println("Input: 4/2       →", validateMonth("4/2"))
	fmt.Println("Input: 5-8       →", validateMonth("5-8"))
	fmt.Println("Input: JAN,MAR   →", validateMonth("JAN,MAR"))
	fmt.Println("Input: JUL-OCT   →", validateMonth("JUL-OCT"))

	fmt.Println("-- Day of Week --")
	fmt.Println("Input: MON-WED   →", validateDayOfWeek("MON-WED"))
	fmt.Println("Input: 1-5       →", validateDayOfWeek("1-5"))
	fmt.Println("Input: 1,3,5     →", validateDayOfWeek("1,3,5"))
	fmt.Println("Input: 0         →", validateDayOfWeek("0"))

}

func validateMinute(minuteField string) string {
	if strings.Contains(minuteField, "*") {
		return "Every minute"
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
		return "Every " + strconv.Itoa(stepInterval) + " minutes, starting at " + strconv.Itoa(stepStart) + " minutes past the hour"
	} else if strings.Contains(minuteField, "-") {
		rangeStart, rangeEnd := handleRange(minuteField)
		return "Minutes " + strconv.Itoa(rangeStart) + " through " + strconv.Itoa(rangeEnd) + " past the hour"
	}
	return "At " + minuteField + " minutes past the hour"
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
		return "Every " + strconv.Itoa(stepInterval) + " hours, starting at " + convert24To12(strconv.Itoa(stepStart))
	} else if strings.Contains(hourField, "-") {
		rangeStart, rangeEnd := handleRange(hourField)
		return "between " + convert24To12(strconv.Itoa(rangeStart)) + " and " + convert24To12(strconv.Itoa(rangeEnd))
	}
	return "between " + convert24To12(hourField) + " and " + convert24To12(hourField)
}

func validateDayOfMonth(dayOfMonthField string) string {
	if strings.Contains(dayOfMonthField, ",") {
		values := strings.Split(dayOfMonthField, ",")
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
	} else if strings.Contains(dayOfMonthField, "/") {
		stepStart, stepInterval := handleStep(dayOfMonthField)
		if stepStart == 1 {
			return "Every " + strconv.Itoa(stepInterval) + " days"
		}
		return "Every " + strconv.Itoa(stepInterval) + " days, starting on day " + strconv.Itoa(stepStart) + " of the month"
	} else if strings.Contains(dayOfMonthField, "-") {
		rangeStart, rangeEnd := handleRange(dayOfMonthField)
		return "between day " + strconv.Itoa(rangeStart) + " and " + strconv.Itoa(rangeEnd) + " of the month"
	}

	// else if dayOfMonthField == "L" {
	// 	return "on the last day of the month"
	// }
	return "on day " + dayOfMonthField + " of the month"
}

func validateMonth(monthField string) string {
	months := []string{"JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE",
		"JULY", "AUGUST", "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER"}

	if strings.Contains(monthField, ",") {
		values := strings.Split(monthField, ",")
		str := ""

		for index := range values {
			values[index] = months[normalizeValue(values[index])-1]
		}

		for index, value := range values {
			str += value
			if index == len(values)-2 {
				str += " and " + values[index+1]
				break
			}
			str += ", "
		}
		return "Only in " + str
	} else if strings.Contains(monthField, "/") {
		stepStart, stepInterval := handleStep(monthField)
		if stepStart == 1 {
			return "Every " + strconv.Itoa(stepInterval) + " months"
		}
		return "Every " + strconv.Itoa(stepInterval) + " months, from " + months[stepStart-1] + " through DECEMBER"
	} else if strings.Contains(monthField, "-") {
		rangeStart, rangeEnd := handleRange(monthField)
		return months[rangeStart-1] + " through " + months[rangeEnd-1]
	}
	return "Only in " + months[normalizeValue(monthField)-1]
}

func validateDayOfWeek(dayOfWeekField string) string {
	weeks := []string{"SUNDAY", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY"}

	if strings.Contains(dayOfWeekField, ",") {
		values := strings.Split(dayOfWeekField, ",")
		str := ""

		for index := range values {
			values[index] = weeks[normalizeValue(values[index])]
		}

		for index, value := range values {
			str += value
			if index == len(values)-2 {
				str += " and " + values[index+1]
				break
			}
			str += ", "
		}
		return "Only on " + str
	} else if strings.Contains(dayOfWeekField, "/") {
		stepStart, stepInterval := handleStep(dayOfWeekField)
		return "Every " + strconv.Itoa(stepInterval) + " days of the week, " + weeks[stepStart] + " through DECEMBER"
	} else if strings.Contains(dayOfWeekField, "-") {
		rangeStart, rangeEnd := handleRange(dayOfWeekField)
		return weeks[rangeStart] + " through " + weeks[rangeEnd]
	}
	return "Only on " + weeks[normalizeValue(dayOfWeekField)]
}

func convert24To12(hr string) string {
	i, _ := strconv.Atoi(hr)
	if i == 0 {
		return "12:00 AM"
	} else if i == 12 {
		return "12:00 PM"
	} else if i < 12 {
		return fmt.Sprintf("%d:00 AM", i)
	}
	return fmt.Sprintf("%d:00 PM", i-12)
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
	weeks := []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}

	normalize, err := strconv.Atoi(value)
	if err != nil {
		for index, month := range months {
			if month == value {
				return index + 1
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
