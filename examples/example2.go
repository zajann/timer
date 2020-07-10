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

	for i := 0; i < 10; i++ {
		fmt.Println(timer.GetTimeFormat("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}
