package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(2000 * time.Millisecond)
	}
}

func hi(msg string) {
	time.Sleep(8000 * time.Millisecond)
	fmt.Println(msg)
}

func main() {
	start := time.Now()
	go hi("고루틴2")                // 새 고루틴에서 실행
	go say("고루틴1")               // 새 고루틴에서 실행
	time.Sleep(10 * time.Second) // 메인 고루틴 대기
	fmt.Println("전체 실행 시간 : ", time.Since(start))
}
