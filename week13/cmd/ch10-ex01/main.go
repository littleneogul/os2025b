package main

import (
	"fmt"
	"log"
	"week13/cmd/pkg/calendar"
)

func main() {
	today := calendar.Date{}
	// today.year = 2025
	// today.month = 10
	// today.day = 20

	err := today.SetYear(2025)
	if err != nil {
		log.Fatal(err)
	}

	err = today.SetMonth(11)
	if err != nil {
		log.Fatal(err)
	}

	err = today.SetDay(24)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(today.Year(), "년 ", today.Month(), "월 ", today.Day(), "일")
}
