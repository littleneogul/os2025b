package main

import (
	"fmt"
	"reflect"

	"github.com/headfirstgo/magazine"
)

type Gallons float64
type Liters float64

func main() {
	var carFuel Gallons
	var busFuel Liters

	carFuel = Gallons(100.0)

	busFuel = Liters(100.0)

	fmt.Println(reflect.TypeOf(carFuel), reflect.TypeOf(busFuel), carFuel, busFuel)

	var s1 magazine.Subscriber
	var e1 magazine.Employee
	s1.Name = "최인하"
	e1.Name = "이인하"
	e1.Salary = 10000000
	e1.Address.City = "인천"
	s1.Address.City = "서울"
	fmt.Println(e1.Name, e1.Salary)
	fmt.Println(s1.Name, s1.Address.City)
}
