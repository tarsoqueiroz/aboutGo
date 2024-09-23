package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	fmt.Println("Starting timeLapse app\n")

	fmt.Println("... pause for 2 seconds.\n")
	// pause program execution for 2 seconds
	time.Sleep(2 * time.Second)

	finishTime := time.Now()
	sinceTime := time.Since(startTime)
	durationTime := finishTime.Sub(startTime)

	fmt.Println("Start  time.:", startTime)
	fmt.Println("Finish time.:", time.Now())
	fmt.Println("Since.......:", sinceTime)
	fmt.Println("Duration....:", durationTime)
	fmt.Println("\nbye...")
}
