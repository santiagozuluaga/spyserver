package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GetTime() string {

	t := time.Now()
	year, month, day := t.Date()
	hour, min, sec := t.Clock()

	date := fmt.Sprintf("%d/%d/%d,%d:%d:%d", year, int(month), day, hour, min, sec)

	return date
}

func UpdateDate(current string) (bool, string) {

	newDate := GetTime()

	splitCurrent := strings.Split(current, ",")
	splitNew := strings.Split(newDate, ",")

	if splitCurrent[0] != splitNew[0] {

		return true, newDate
	} else {

		hourCurrent, _ := strconv.Atoi(strings.Split(splitCurrent[1], ":")[0])
		hourNew, _ := strconv.Atoi(strings.Split(splitNew[1], ":")[0])
		minCurrent, _ := strconv.Atoi(strings.Split(splitCurrent[1], ":")[1])
		minNew, _ := strconv.Atoi(strings.Split(splitNew[1], ":")[1])

		if hourCurrent+1 < hourNew || hourCurrent+1 == hourNew && minCurrent <= minNew {

			return true, newDate
		} else {

			return false, current
		}
	}
}
