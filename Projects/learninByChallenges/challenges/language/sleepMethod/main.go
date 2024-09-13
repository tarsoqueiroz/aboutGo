package main

import (
	"fmt"
	"time"
)

func pause2seconds() {

	// Hi
	fmt.Printf("****************************\n")
	fmt.Printf("***   Pause 2 seconds    ***\n")
	fmt.Printf("****************************\n\n")

	// prints message before sleep
	fmt.Println("Executing code before sleep:", time.Now().Format(time.RFC850))

	// pause program execution for 2 seconds
	time.Sleep(2 * time.Second)

	// prints message after sleep
	fmt.Println("Executing code after  sleep:", time.Now().Format(time.RFC850))

	// Bye
	fmt.Printf("\n****************************\n\n")

}

func pauseVariable() {

	// Hi
	fmt.Printf("****************************\n")
	fmt.Printf("***    Pause variable    ***\n")
	fmt.Printf("****************************\n\n")

	// prints message before sleep looping
	fmt.Printf("Executing code before sleep: %s\n\n", time.Now().Format(time.RFC850))

	// for loop that will run 5 times
	for nIdx := 0; nIdx < 5; nIdx++ {
		// prints message before each sleep
		fmt.Printf("(%02d) Executing code in loop: %s\n", nIdx, time.Now().Format(time.RFC850))

		// pauses program execution for a duration that increases by one second for each iteration of the loop
		time.Sleep(time.Duration(nIdx) * time.Second)

		// prints message after each sleep
		fmt.Printf("(%02d) Executed  code in loop: %s\n\n", nIdx, time.Now().Format(time.RFC850))
	}

	// prints message after sleep
	fmt.Println("Executing code after  sleep:", time.Now().Format(time.RFC850))

	// Bye
	fmt.Printf("\n****************************\n\n")

}

func timerVariable() {

	// Hi
	fmt.Printf("****************************\n")
	fmt.Printf("***    Timer variable    ***\n")
	fmt.Printf("****************************\n\n")

	// prints message before sleep looping
	fmt.Printf("Executing code before timer: %s\n\n", time.Now().Format(time.RFC850))

	// for loop that will run 5 times
	for nIdx := 0; nIdx < 5; nIdx++ {
		// prints message before each sleep
		fmt.Printf("(%02d) Executing code in loop: %s\n", nIdx, time.Now().Format(time.RFC850))

		// creates a timer that will fire after nIdx seconds
		timer := time.NewTimer(time.Duration(nIdx) * time.Second)
		// waits for the timer to fire
		<-timer.C

		// prints message after each sleep
		fmt.Printf("(%02d) Executed  code in loop: %s\n\n", nIdx, time.Now().Format(time.RFC850))
	}

	// prints message after sleep
	fmt.Println("Executing code after  timer:", time.Now().Format(time.RFC850))

	// Bye
	fmt.Printf("\n****************************\n\n")

}

func main() {

	pause2seconds()

	pauseVariable()

	timerVariable()

}
