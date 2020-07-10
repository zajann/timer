package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zajann/timer"
)

func main() {
	if err := timer.Init(0); err != nil {
		log.Fatal(err)
	}

	fmt.Println(timer.GetYear())
	fmt.Println(timer.GetMonth())
	fmt.Println(timer.GetMonthString())
	fmt.Println(timer.GetDay())
	fmt.Println(timer.GetHour())
	fmt.Println(timer.GetMinute())
	fmt.Println(timer.GetSecond())
	fmt.Println(timer.GetNanoSecond())
	fmt.Println(timer.GetUnix())
	fmt.Println(timer.GetUnixNano())

	fmt.Println(timer.GetTimeFormat(time.RFC1123))
	fmt.Println(timer.GetTimeFormat("2006-01-02 15:04:05"))

	start := timer.GetTime()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond)
	}
	elapsed := time.Since(start)

	fmt.Printf("Elapsed Time: %s\n", elapsed)
}
