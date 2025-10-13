package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	log.Fatal(err)
	fmt.Println(i, err)
}
