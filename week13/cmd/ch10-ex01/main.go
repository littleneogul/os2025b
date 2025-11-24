package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/calendar"
)

func main() {
	event := calendar.Event{}

	err := event.SetYear(2025)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetMonth(11)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(24)
	if err != nil {
		log.Fatal(err)
	}

	event.SetTitle("생년월일")

	fmt.Println(event.Title())
	fmt.Println(event.Year(), "년 ", event.Month(), "월", event.Day(), "일")

	today := calendar.Date{}
	// today.year = 2025
	// today.month = 10
	// today.day = 20

	err = today.SetYear(2025)
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
