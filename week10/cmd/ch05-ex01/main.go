package main

import "fmt"

func main() {
	var arrayBool [3]bool = [3]bool{true, false, true}
	var arrayInt [3]int

	fmt.Println(arrayBool[1])
	arrayInt[1] = 2
	fmt.Println(arrayInt[1])
}
