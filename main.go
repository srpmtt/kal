package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseMonth(mStr string) (time.Month, error) {
	mStr = strings.ToLower(mStr)
	months := map[string]time.Month{
		"january":   time.January,
		"february":  time.February,
		"march":     time.March,
		"april":     time.April,
		"may":       time.May,
		"june":      time.June,
		"july":      time.July,
		"august":    time.August,
		"september": time.September,
		"october":   time.October,
		"november":  time.November,
		"december":  time.December,
	}

	if month, found := months[mStr]; found {
		return month, nil
	}

	shortMonths := map[string]time.Month{
		"jan": time.January,
		"feb": time.February,
		"mar": time.March,
		"apr": time.April,
		"may": time.May,
		"jun": time.June,
		"jul": time.July,
		"aug": time.August,
		"sep": time.September,
		"oct": time.October,
		"nov": time.November,
		"dec": time.December,
	}

	if month, found := shortMonths[mStr]; found {
		return month, nil
	}

	monthNum, err := strconv.Atoi(mStr)
	if err == nil && monthNum >= 1 && monthNum <= 12 {
		return time.Month(monthNum), nil
	}

	return 0, fmt.Errorf("invalid month: %s", mStr)
}

func daysInMonth(year int, month time.Month) int {
	lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local)
	return lastDay.Day()
}

func printHeader(m time.Month, year int) {
	month := m.String()
	fmt.Printf("    %s %d\n", month, year)
	fmt.Println("Mo Tu We Th Fr Sa Su")
}

func printMonth(year int, month time.Month) {
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local).Weekday()
	adjustedFirstDay := (int(firstDay) + 6) % 7
	daysInMonth := daysInMonth(year, month)

	for i := 0; i < adjustedFirstDay; i++ {
		fmt.Print("   ")
	}

	for day := 1; day <= daysInMonth; day++ {
		if (adjustedFirstDay+day)%7 == 6 || (adjustedFirstDay+day)%7 == 0 {
			fmt.Printf("\033[31m")
		}

		if day < 10 {
			fmt.Printf(" %d ", day)
		} else {
			fmt.Printf("%d ", day)
		}

		if (adjustedFirstDay+day)%7 == 6 || (adjustedFirstDay+day)%7 == 0 {
			fmt.Printf("\033[0m")
		}

		if (adjustedFirstDay+day)%7 == 0 {
			fmt.Println()
		}
	}

	if (adjustedFirstDay+daysInMonth)%7 != 0 {
		fmt.Println()
	}
}

func printYear(year int) {
	for row := 0; row < 3; row++ {
		for col := 0; col < 4; col++ {
			month := time.Month(row*4 + col + 1)
			printHeader(month, year)
			printMonth(year, month)
			if col < 3 {
				fmt.Print("\n")
			}
		}
		fmt.Println()
	}
}

func main() {
	currentTime := time.Now()
	month := currentTime.Month()
	year := currentTime.Year()

	if len(os.Args) > 1 {
		if len(os.Args) == 2 {
			year, _ = strconv.Atoi(os.Args[1])
		} else if len(os.Args) == 3 {
			monthStr := os.Args[1]
			year, _ = strconv.Atoi(os.Args[2])
			var err error
			month, err = parseMonth(monthStr)

			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	if len(os.Args) == 2 {
		printYear(year)
	} else {
		printHeader(month, year)
		printMonth(year, month)
	}
}
