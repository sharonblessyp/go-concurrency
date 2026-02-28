package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

/*
Use when tasks can run independently or in parallel (e.g., send email, process jobs)
*/
func main() {
	go printNumbers()

	time.Sleep(1 * time.Second)
	fmt.Println("Main finished")
}
