package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("I:", i)
		time.Sleep(time.Second)
	}

	fmt.Println(i)
}
