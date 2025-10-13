package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var month time.Month = now.Month()
	var day int = now.Day()
	fmt.Println(month, day)

	univ := "Go$ Inha$"
	changer := strings.NewReplacer("$", "!")
	changed := changer.Replace(univ)
	fmt.Println(changed)
}
