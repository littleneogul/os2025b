// average calculates the average of several numbers.
package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	weights, err := datafile.GetFloats("meatWeight.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0 // float64로 명시 혹은 sum := 0.0
	for _, number := range weights {
		sum += number
	}
	weeks := float64(len(weights))
	fmt.Println("평균 고기 소비량 : ", sum/weeks)
}
