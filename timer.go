package timer

import (
	"errors"
	"time"
)

var (
	initTimer bool
	t         time.Time
	tYear     int
	tMonth    time.Month
	tDay      int
	tHour     int
	tMinute   int
	tSec      int
	tNanoSec  int
	tUnix     int64
	tUnixNano int64
	sleep     time.Duration
)

func Init(d time.Duration) error {
	if initTimer {
		return errors.New("already init timer")
	}
	initTimer = true
	if d == 0 { // if d is zero, set default value (100 millisecond)
		sleep = time.Millisecond * 100
	} else {
		sleep = d
	}
	updateTime()
	go updateTimeGoroutine()
	return nil
}

func GetTimeFormat(layout string) string {
	return t.Format(layout)
}

func GetTime() time.Time {
	return t
}

func GetYear() int {
	return tYear
}

func GetMonth() int {
	return int(tMonth)
}

func GetMonthString() string {
	return tMonth.String()
}

func GetDay() int {
	return tDay
}

func GetHour() int {
	return tHour
}

func GetMinute() int {
	return tMinute
}

func GetSecond() int {
	return tSec
}

func GetNanoSecond() int {
	return tNanoSec
}

func GetUnix() int64 {
	return tUnix
}

func GetUnixNano() int64 {
	return tUnixNano
}

func Stop() {
	if !initTimer {
		initTimer = false
	}
}

func updateTime() {
	t = time.Now()

	tYear = t.Year()
	tMonth = t.Month()
	tDay = t.Day()
	tHour = t.Hour()
	tMinute = t.Minute()
	tSec = t.Second()
	tNanoSec = t.Nanosecond()
	tUnix = t.Unix()
	tUnixNano = t.UnixNano()
}

func updateTimeGoroutine() {
	for initTimer {
		updateTime()
		time.Sleep(sleep)
	}
}
