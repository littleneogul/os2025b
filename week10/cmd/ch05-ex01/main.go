package main

import (
	"fmt"
	"reflect"
)

func main() {
	arrayBool := [2]bool{true, false}
	arrayInt := [3]int{-9, 11, 7}

	for i := 0; i < len(arrayInt); i++ {
		fmt.Println(i, arrayBool[i])
		fmt.Println(i, arrayInt[i])
	}

	fmt.Printf("%#v\n", arrayBool)
	fmt.Printf("%#v\n", arrayInt)
	fmt.Println(reflect.TypeOf(arrayInt))
}
