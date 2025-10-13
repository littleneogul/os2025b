package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)                // trim 함수
	score, err := strconv.ParseFloat(i, 64) // 정리된 문자열을 실수타입으로 변환
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if score >= 60 {
		status = "Pass"
	} else {
		status = "Fail"
	}
	fmt.Println(score, "is", status)
}
