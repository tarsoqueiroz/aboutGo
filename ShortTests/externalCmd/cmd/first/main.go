package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("ls", "-ltr")

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)

}
