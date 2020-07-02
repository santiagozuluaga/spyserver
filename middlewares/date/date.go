package date

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

var MONTHS = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December"}

func GetTime() (int, int, int, int, int, int) {

	t := time.Now()
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	var monthNumber int

	for i := 0; i < 12; i++ {
		if MONTHS[i] == month.String() {
			monthNumber = i + 1
		}
	}

	return year, monthNumber, day, hour, min, sec
}

func CreateDate() string {

	year, month, day, hour, min, sec := GetTime()

	date := fmt.Sprintf("%d/%d/%d,%d:%d:%d", year, month, day, hour, min, sec)

	return date
}

func UpdateDate(current string) (bool, string) {

	flag := true
	newYear, newMonth, newDay, newHour, newMin, newSec := GetTime()
	var dateInt []int
	var clockInt []int

	split := strings.Split(current, ",")
	date := strings.Split(split[0], "/")
	clock := strings.Split(split[1], ":")

	for i := 0; i < len(date); i++ {

		str, err := strconv.Atoi(date[i])
		if err != nil {

			log.Println(err)
		}

		dateInt = append(dateInt, str)
	}

	for i := 0; i < len(clock); i++ {

		str, err := strconv.Atoi(clock[i])
		if err != nil {

			log.Println(err)
		}

		clockInt = append(clockInt, str)
	}

	if dateInt[0] < newYear || dateInt[1] < newMonth || dateInt[2] != newDay {

		new := fmt.Sprintf("%d/%d/%d,%d:%d:%d", newYear, newMonth, newDay, newHour, newMin, newSec)
		return flag, new
	} else {

		if clockInt[0] != newHour {

			new := fmt.Sprintf("%d/%d/%d,%d:%d:%d", newYear, newMonth, newDay, newHour, newMin, newSec)
			return flag, new

		} else {

			flag = false
			return flag, current
		}

	}
}
