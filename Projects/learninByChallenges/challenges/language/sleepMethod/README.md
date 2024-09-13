# Golang: Practical Cases to Use the Golang Sleep Method

> `https://dev.to/free_coder/golang-practical-cases-to-use-the-golang-sleep-method-74p`

When it comes to concurrent programming in Go, you may need to handle a Golang sleep or pause program execution for a specific amount of time. To accomplish this, Go provides the time package with a Sleep() method. In this guide, we'll show you how to use the Golang sleep() method in deep with examples and comments, and cover some related topics.

## Starting

```sh
go mod init sleepmethod

touch main.go
```

## Using the Golang Sleep Method

- `func pause2seconds()`: This function pauses for 2 seconds before printing the final message.

```go
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
```

## Golang Sleeping & Pausing for a Variable Duration

- `func pauseVariable()`: This function executes the code inside the loop and pauses for a duration that increases by one second for each iteration of the loop.

```go
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
```

## Golang Sleep Using Timers

- `func timerVariable()`: In addition to the Golang sleep method, the time package in Go provides other useful tools for working with time. One of them is the Timer struct, which you can use to schedule an event to occur after a certain duration.

```go
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
```

## That's all

...folks!!!
